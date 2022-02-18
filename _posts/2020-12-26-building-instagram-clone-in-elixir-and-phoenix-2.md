---
layout: post
title: Building an Instagram clone in Elixir and Phoenix, Part 2
date: '2020-10-27'
summary: 
icon: icons/phoenix.png
comments: true
category:
  - elixir
  - phoenix
---

Welcome back to the series. Today we're going to add users to Picton. These users will be able to create a network of followers, as weel as to share moments with their network.

Let's start by describing what a user will see while interacting with our application, and we'll assume we already have all the functionality implemented. That way we can write our tests for what the application should behave like, and then make those tests pass.

So, as a non-logged in user, when I visit Picton's homepage, I won't see any moments posted by any of the users. Instead, I will see a Log In form where I can put my username and password and if my credentials are correct, after submitting the form I'll see moments posted by others. On the other hand, if my credentials are wrong, I'll see an error message. 

Let's describe this with a test. Add this in `test/picton_web/features/authentication.exs`:

```elixir
defmodule PictonWeb.AuthenticationTest do
  use PictonWeb.IntegrationCase

  @username "janedoe"
  @password "v3rys3cur3p4ssw0rd"

  describe "Authentication" do
    feature "User logs in with valid credentials", %{session: session} do
      session
      |> visit("/")
      |> assert_has("Log In")
      |> assert_has("Sign Up")
      |> assert_has(Query.css("form.login"))
      |> fill_in(Query.fillable_field("user[username]"), with: @username)
      |> fill_in(Query.fillable_field("user[password]"), with: @password)
      |> click(Query.css("form.login [type='submit']"))
      |> assert_has(Query.css(".new-moment--caption"))
      |> assert_has("Log out")
    end

    feature "User tries to log in with invalid credentials", %{session: session} do
      session
      |> visit("/")
      |> assert_has("Log In")
      |> assert_has("Sign Up")
      |> assert_has(Query.css("form.login"))
      |> fill_in(Query.fillable_field("user[username]"), with: @username)
      |> fill_in(Query.fillable_field("user[password]"), with: "hehe.Hacked")
      |> click(Query.css("form.login [type='submit']"))
      |> assert_has("Your username or password is incorrect.")
    end
  end
end
```

So we expect the home page to have a Log in form, as well as some way to let us sign up. After we log in with valid credentials, we should see the new Moment form, as well as a way to log out. If our credentials are wrong, we should see an error message. 

Feel free to run the test, but since we have no login form it's obvious that it will fail.

## Authenticating Users

If you came from the Rails world, you're definitely familiar with Devise. We have an equivalent of the bloated authentication gem for Phoenix, and it's called [Pow](https://github.com/danschultzer/pow). Add it in your `mix.exs`:

```elixir
defp deps do
  [
    # ...
    {:pow, "~> 1.0.21"}
  ]
end
```

Now, install the dependencies and then install Pow files:

```bash
$ mix deps.get
$ mix pow.install
```

When the installation is done, follow the instructions printed in the console to add the configurations, the plug and the routes necessary for Pow to work. 

Update your `lib/picton_web/router.ex` to this:

```elixir
pipeline :protected do
  plug Pow.Plug.RequireAuthenticated,
        error_handler: Pow.Phoenix.PlugErrorHandler
end

scope "/" do
  pipe_through :browser

  pow_routes()
end

scope "/", PictonWeb do
  pipe_through [:browser, :protected]

  get "/", MomentController, :index
  resources "/moments", MomentController, except: [:new]
end
```

Don't worry why we need two different `scope "/"` for the moment, we'll come to that later. If you didn't notice, Pow added a User schema in `lib/picton/users/user.ex`, as well as a migration in `priv/repo/migrations/XXX_create_users.exs`. We'll update the migration to look like this:

```elixir
defmodule Picton.Repo.Migrations.CreateUsers do
  use Ecto.Migration

  def change do
    create table(:users) do
      add :email, :string, null: false
      add :username, :string, null: false
      add :password_hash, :string

      timestamps()
    end

    create unique_index(:users, [:email])
    create unique_index(:users, [:username])
  end
end
```

We added a unique `username` field which we'll be using to authenticate users. Go ahead and run the migration. If you don't remember the Ecto task, you can refer to [the documentation](https://hexdocs.pm/ecto_sql/Ecto.Migration.html#module-mix-tasks). 

Now we need to add this field to our schema as well, otherwise Ecto won't know how to handle the username:

```elixir
defmodule Picton.Users.User do
  use Ecto.Schema
  use Pow.Ecto.Schema
  import Ecto.Changeset

  schema "users" do
    pow_user_fields()
    field :username, :string

    timestamps()
  end

  def changeset(user, attrs) do
    user
    |> pow_changeset(attrs)
    |> cast(attrs, [:username])
    |> validate_required([:username])
  end
end
```

If you run the server and visit [http://localhost:4000/registration/new](http://localhost:4000/registration/new), you'll see a registration form, asking you for an email and a password. Try filling out the form and you'll see it returns an error when you submit it. That's because we marked the username field as required in the changeset above. We will now update the registration form to allow us enter a username as well.

Start by setting updating the Pow config in `config/config.exs` to this:

```elixir
config :picton, :pow,
  user: Picton.Users.User,
  repo: Picton.Repo,
  web_module: PictonWeb # Add this
```

Then, run this:

```bash
$ mix pow.phoenix.gen.templates
```

This will generate some views and templates, and we need to add this into `lib/picton_web/templates/pow/registration/new.html.eex`:

```eex
<!-- Email fields -->

<%= label f, :username %>
<%= text_input f, :username %>
<%= error_tag f, :username %>

<!-- Password fields -->
```

Now try filling the form again and see what happens when you submit it. It works fine, right?

## Protected routes

When we added Pow routes to `lib/picton_web/router.ex`, we created a `:protected` pipeline. These protected routes are the routes which can be accessed only after the user is logged in. We'll make all the moment-related routes protected, and we'll leave the Log In/Sign Up routes unprotected.

We can check which routes Pow added for us:

```bash
$ mix phx.routes | grep Pow

     pow_session_path  GET     /session/new          Pow.Phoenix.SessionController :new
     pow_session_path  POST    /session              Pow.Phoenix.SessionController :create
     pow_session_path  DELETE  /session              Pow.Phoenix.SessionController :delete
pow_registration_path  GET     /registration/edit    Pow.Phoenix.RegistrationController :edit
pow_registration_path  GET     /registration/new     Pow.Phoenix.RegistrationController :new
pow_registration_path  POST    /registration         Pow.Phoenix.RegistrationController :create
pow_registration_path  PATCH   /registration         Pow.Phoenix.RegistrationController :update
                       PUT     /registration         Pow.Phoenix.RegistrationController :update
pow_registration_path  DELETE  /registration         Pow.Phoenix.RegistrationController :delete
```

You can see that Pow renders the Log In form in a `/session/new` page; we want it in our root. So let's make this happen. We have a `PictonWeb.PageController` which was generated with our application. So let's use its index page as our application's landing page. I found [this Unsplash image](https://unsplash.com/photos/xDPZ3xTEh2g) by [Anthony Vela](https://unsplash.com/@anthonyvela) which I thought matched Picton. Now open `lib/picton_web/templates/page/index.html.eex` and update it to the following:

```eex

```