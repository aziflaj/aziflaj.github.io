---
layout: post
title: From ActiveRecord to Ecto
date: '2020-10-26'
summary: Migrating an old Rails app into an Elixir one
icon: icons/phoenix.png
comments: true
category:
  - elixir
  - phoenix
  - rails
  - ruby 
---

Hello dear web surfer. Are you too annoyed by all the articles in the internet about Ruby being a dead language? Don't you wish there was a new language, tailored to Ruby developers, which is not dead yet? Look no further my friend, Elixir is here.

[Elixir](https://elixir-lang.org/) is this niche language developed by [Jose Valim](https://twitter.com/josevalim), who was one of the core contributors in Rails. It runs on [Erlang](https://www.erlang.org/) VM (BEAM) and shares a lot of features and tools with Erlang, which has shined in (soft-)real-time high availability systems. The whole infrastructure of WhatsApp is built upon Erlang, and other companies like Heroku and Discord are using Elixir.

The way Elixir achieves its high availability is a topic for another time. Learning all the OTP jibber-jabber will take some time, but you don't actually need to know everything regarding OTP in order to build applications. If you're already used to Rails and want to give Elixir a try, [Phoenix](https://www.phoenixframework.org/) is right up your alley.

Here I want to show you some code samples from an old Rails application of mine, and how I'm rewriting the same code in Elixir using Ecto instead of Rails' ActiveRecord. "What is Ecto?" you might ask. And to that I say "good question". [Ecto](https://github.com/elixir-ecto/ecto) is, in the broad sense of the word, the gateway to your application's data. You will use Ecto and its toolkit to change the structure of the database, map data from the DB to your Elixir structs and vice versa, run queries, etc.

Without further ado, let's jump right into it.

## Creating migrations

Ecto's migrations are very similar to those in Rails. You can generate them using mix:

```bash
$ mix ecto.gen.migration create_slots
```

Here's the Ruby code I wrote for this migration:

```ruby
class CreateSlots < ActiveRecord::Migration[6.0]
  def change
    create_table :slots do |t|
      t.datetime :scheduled_at
      t.references :user, null: false, foreign_key: true
      t.text :notes
      t.string :guest_email

      t.timestamps
    end
  end
end
```

And here's the Elixir code:

```elixir
defmodule Skeduler.Repo.Migrations.CreateSlots do
  use Ecto.Migration

  def change do
    create table(:slots) do
      add :scheduled_at, :utc_datetime
      add :user_id, references(:users), null: false
      add :notes, :text
      add :guest_email, :string

      timestamps()
    end
  end
end
```

If you don't already know this, Elixir is a functional language. Instead of classes we use modules, defined with `defmodule`. Instead of inheriting from `ActiveRecord::Migration` we `use Ecto.Migration`, which exposes all the migration-related functions into our module. The code is very similar to Rails' migrations: this is how Elixir lures Ruby developers in, only to trap them inside those long-running, self-isolated processes that aren't really processes (they're [actors](https://en.wikipedia.org/wiki/Actor_model), but more on that some other time).

Running migrations is as easy as in Rails too:

```bash
$ mix ecto.migrate
```

## Modeling your Data

Here's my old Rails model:

```ruby
class Slot < ApplicationRecord
  belongs_to :user, dependent: :destroy, required: true

  validates :guest_email, 
            allow_blank: true,
            format: { with: /\A[A-Z0-9._%a-z\-]+@(?:[A-Z0-9a-z\-]+\.)+[A-Za-z]{2,4}\z/ }

  validate :scheduled_in_the_future

  def booked?
    guest_email.present?
  end

  private

  def scheduled_in_the_future
    seconds_params = { sec: 0 }
    return if scheduled_at.change(seconds_params) >= 5.minutes.from_now.change(seconds_params)

    errors.add(:scheduled_at, 'should be at least 5 minutes from now')
  end
end
```

And here's my new Ecto model:

```elixir
defmodule Skeduler.Slot do
  use Ecto.Schema
  use Timex
  import Ecto.Changeset
  alias __MODULE__

  schema "slots" do
    field :scheduled_at, :utc_datetime
    belongs_to :user, Skeduler.User
    field :notes, :text
    field :guest_email, :string

    timestamps()
  end

  def changeset(%Slot{} = slot, attrs) do
    slot
    |> cast(attrs, [:scheduled_at, :user_id, :notes, :guest_email])
    |> validate_required([:guest_email, :scheduled_at, :user_id])
    |> validate_format(:guest_email, ~r/[A-Z0-9._%a-z\-]+@(?:[A-Z0-9a-z\-]+\.)+[A-Za-z]{2,4}/)
    |> validate_scheduled_in_the_future()
  end

  defp validate_scheduled_in_the_future(changeset) do
    validate_change(changeset, :scheduled_at, fn (_, scheduled_time) ->
      valid_time = 
        Timex.now
        |> Timex.shift(minutes: 5)
        |> Timex.before?(scheduled_time)
        |> case do
             true -> []
             false -> [scheduled_at: "should be at least 5 minutes from now"]
           end
    end)
  end
end
```

Ecto uses the `schema` block to map the data from your DB to a `%Skeduler.Slot{}` struct and the other way around. Also, it needs the `changeset/2` function to be implemented, so we can check if our data is valid. I'm using [Timex](https://github.com/bitwalker/timex) here to handle Date and Time calculations; Elixir doesn't have the syntatic `5.minutes.from_now` sugar that comes with Rails. What Elixir has is the funny way of chaining function calls, also known as pipelining. See the `|>` operator which links `Timex.now` to some function calls and then to a `case`?

Instead of `Timex.now |> Timex.shift(minutes: 5)`, you can write `Timex.shift(Timex.now, minutes: 5)`; both of these yield the same result. But since Elixir comes with a pipe operator, I'm gonna overuse it when I can. 

## Querying your Data

In Rails, you've already used the built-in ActiveRecord methods and if those methods failed you, you probably jumped into the Arel train. Sometimes, to keep it clean, we put these querying functions into reusable modules and call them Query Objects. [You can see here](https://github.com/aziflaj/skeduler/blob/master/app/queries) my Rails implementation of two simple query objects. 

Ecto does it a bit different. If you ever used C#, you're probably familiar with LINQ. Meet `Ecto.Query`:

```elixir
defmodule Skeduler.SlotQueries do
  import Ecto.Query
  alias Skeduler.{User, Slot}

  def free_slots_for_user(%User{id: user_id}) do
    from s in Slot,
      where: (s.user_id == ^user_id),
      where: (is_nil(s.guest_email) or s.guest_email == ""),
      order_by: s.scheduled_at
  end

  def upcoming_slots_for_user(%User{id: user_id}) do
    from s in Slot,
      where: (s.user_id == ^user_id) and
             not (is_nil(s.guest_email) or s.guest_email == ""),
      order_by: s.scheduled_at
  end
end
```

`Ecto.Query` makes it look like you're writting SQL queries in an Elixir-like syntax, but it's just smoke and mirrors. You're not writing actual queries, you're using what's called [a keyword list](https://elixir-lang.org/getting-started/keywords-and-maps.html#keyword-lists): an Elixir trickery that looks almost exactly like a Ruby Hash but allows key repetition. That's why you can have multiple `where` "clauses" in the `free_slots_for_user/1` function.

And here's how yoy use the queries above:

```elixir
alias Skeduler.{Repo, SlotQueries}

free_slots = 
  user
  |> SlotQueries.free_slots_for_user()
  |> Repo.all()
```

And that's it for the time being. You can read more about Ecto [in the official docs](https://hexdocs.pm/ecto/Ecto.html) or by jumping head first into [its source code](https://github.com/elixir-ecto/ecto). 

_If you have anything to add, any suggestion or question, comment below and I'll get back at ya_