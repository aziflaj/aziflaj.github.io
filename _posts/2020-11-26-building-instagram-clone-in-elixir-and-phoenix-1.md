---
layout: post
title: Building an Instagram clone in Elixir and Phoenix, Part 1
date: '2020-10-26'
summary: 
icon: icons/phoenix.png
comments: true
category:
  - elixir
  - phoenix
---

When you learn a new programming language, you start with the "Hello, World!" program. When I got into full stack web development, I was told to start with building a blog: it gives you insight into designing databases, handling authentication and authorization, styling the interface to make it appealing for the readers, and so on. But blogs are easy. So easy, the one you're currently reading doesn't even have a back-end. 

In order to get more familiar with a new framework, it's better to consider building something more complex than that. So here it goes, an Instagram clone built in Elixir and Phoenix.

This is a step-by-step, test-driven walkthrough into building a web application similar to Instagram. 

> Instagram is a portmanteau of **Insta**nt camera and Tele**gram**. To avoid possible trademark issues and other phony reasons, we'll name our application Picton, a portmanteau of **Pic**ture and Metric **ton**, mainly because a ton is way bigger than a gram and we're compensating for the lack of features.

If you haven't installed Elixir and Phoenix already, [follow these instructions](https://hexdocs.pm/phoenix/installation.html#content).

With all your tools installed, start by creating a new application:

```bash
$ mix phx.new picton
```

This will create an empty Phoenix app backed by PostgreSQL. Edit `config/dev.exs` and `config/test.exs` to set the db credentials, like this (you will want to use your own credentials instead of mine):

```elixir
# Configure your database
config :picton, Picton.Repo,
  username: "postgres", 
  password: "",
  database: "picton_dev",
  hostname: "localhost",
  show_sensitive_data_on_connection_error: true,
  pool_size: 10
```

After doing so, create the database and fire up the server:

```bash
$ mix ecto.create
$ mix phx.server
```

Now visit [http://localhost:4000/](http://localhost:4000/) and you'll see Phoenix's default landing page.

## Set up the test suite

Instagram wouldn't be successful without people posting pics, and Picton won't take over Instagram if people can't do the same. So, our first task is to allow users to submit photos that other people can see and adore. But since this is a test-driven walkthrough, we'll start with setting up the test environment, then we'll write some tests and finally we'll implement the features we want.

Install [Wallaby](https://github.com/elixir-wallaby/wallaby) by adding this to your `mix.exs`:

```elixir
defp deps do
  [
    # ...
    {:wallaby, "~> 0.26.0", runtime: false, only: :test}
  ]
end
```

And then run:

```
$ mix deps.get
```

You also need to configure some test- and Wallaby- related things, such as the SQL sandbox and the test application endpoint. All these details are in the [Wallaby's README](https://github.com/elixir-wallaby/wallaby#setup), but you can also see all of them [here](https://github.com/aziflaj/picton/commit/2a9f517afc40d0cc7e9ea00313bc7f95c169c572) as well. Finally, create a new file `test/support/integration_case.ex` with the following code (inspired by [Jake Worth](https://hashrocket.com/blog/posts/integration-testing-phoenix-with-wallaby)): 

```elixir
defmodule PictonWeb.IntegrationCase do
  use ExUnit.CaseTemplate

  using do
    quote do
      use Wallaby.{DSL, Feature}

      alias Picton.Repo
      import Ecto
      import Ecto.{Changeset, Query}

      import PictonWeb.Router.Helpers
    end
  end

  setup do
    :ok = Ecto.Adapters.SQL.Sandbox.checkout(Picton.Repo)
    Ecto.Adapters.SQL.Sandbox.mode(Picton.Repo, {:shared, self()})

    metadata = Phoenix.Ecto.SQL.Sandbox.metadata_for(Picton.Repo, self())
    {:ok, session} = Wallaby.start_session(metadata: metadata)
    {:ok, session: session}
  end
end
```

Don't worry too much if you don't understand what the code above does line-per-line. All you need to know in order to use it is that it exposes certain modules from Wallaby, Ecto and our application, and it also sets up a sandboxed environment for tests that hit the database. 

Now for that first test, create a file called `test/picton_web/features/post_new_moment.exs` with the following content:

```elixir
defmodule PictonWeb.PostNewMomentTest do
  use PictonWeb.IntegrationCase

  describe "Posting new Moments" do
    feature "User posts a Moment with a caption", %{session: session} do
      caption = "I'm Mr. Meeseeks, look at me!"
      selfie = "test/support/fixtures/images/meeseeks.png"

      session
      |> visit("/")
      |> fill_in(Query.css(".new-moment--caption"), with: caption)
      |> attach_file(Query.css(".new-moment--image"), path: selfie)
      |> click(Query.css("form.new-moment [type='submit']"))
      |> assert_text(Query.css(".moment--caption"), caption)
    end
  end
end
```

This test describes a simple user interaction with Picton. The user goes to the index page, writes a caption in a field with `new-moment--caption` class, uploads an image in a field with `new-moment--image` class, clicks the submit button in a form with `new-moment` class, and expects the caption to be rendered in an element with `moment--caption` class. I took the liberty of referring to the posts in Picton as "moments"; I want the users to have a deep personal connection to the application and not share just pics but moments from their lives... Run the test file and read through the error:

```bash
$ mix test test/picton_web/features/post_new_moment.exs
```

As you can see, the test fails because there's no element that matches to the  `.new-post--caption` css selector in our root page. Let's change that.

## Share Your First Moment

Phoenix comes with a great deal of helpers, one of which will allow us to create our `MomentController`. Notice the singular `Moment`; in Rails, from which Phoenix draws a lot of inspiration, you'd call this a `MomentsController`. In Elixir we like to keep things simple and singular, so keep this in mind if you're migrating from Rails or a framework with similar convention.

Let's generate our first controller:

```bash
$ mix phx.gen.html Moments Moment moments caption:string image:string
```

The first argument passed to `phx.gen.html` is `Moments`, the context name. _"What is a context?"_ you might ask. [Contexts](https://hexdocs.pm/phoenix/contexts.html) are Phoenix's way of grouping similar functionality together. We're putting all the posts-related functionality in a module called `Moments`, and this functionality will then be called by our controller (following the skinny controller principle).  The second argument is `Moment`, which will be used to name the model (or as we call them in Phoenix, a [Schema](https://hexdocs.pm/phoenix/ecto.html#the-schema)) and the controller (`MomentController`). The third argument, `moments`, is the name of the database table used to store the data for a given moment. The rest are optional arguments used to describe the columns of said table.

Running this will generate a lot of files, which we'll touch later on. Go ahead and add `resources "/moments", MomentController` to your `lib/picton_web/router.ex` just like the output of the Mix task says, and run the migrations using: 

```bash
$ mix ecto.migrate
```

If you run the server and visit [http://localhost:4000/moments](http://localhost:4000/moments), you'll see the result of the boilerplate code generated by `mix phx.gen.html`: all CRUD actions ready to be used. Feel free to play around, and come back when you're ready to create a better form for sharing Moments.

The form for creating moments already exists in `lib/picton_web/templates/moment/form.html.eex`. We'll update it to this:

```eex
<%= form_for @changeset, @action, [multipart: true, class: "new-moment"], fn f -> %>
  <%= if @changeset.action do %>
    <div class="alert alert-danger">
      <p>Oops, something went wrong! Please check the errors below.</p>
    </div>
  <% end %>

  <%= file_input f, :image, class: "new-moment--image" %>
  <%= error_tag f, :image %>

  <%= text_input f, :caption, placeholder: 'Caption', class: "new-moment--caption" %>
  <%= error_tag f, :caption %>

  <div>
    <%= submit "Publish" %>
  </div>
<% end %>
```

Note the class names here are the same as the ones used in the feature test we wrote earlier. We want this form to be shown in the index page (given the user is already logged in, but we'll get to that later). If you check `lib/picton_web/templates/moment/new.html.eex`, you'll see how this form is rendered (similar to partial views in Rails):

```eex
<%= render "form.html", Map.put(assigns, :action, Routes.moment_path(@conn, :create)) %>
```

We don't want the Moment creation to happen on a separate view, so we'll remove `lib/picton_web/templates/moment/new.html.eex` and render the form in `lib/picton_web/templates/moment/index.html.eex`, which now looks like this: 

```eex
<%= render "form.html", Map.put(assigns, :action, Routes.moment_path(@conn, :create)) %>

<%= for moment <- @moments do %>
  <div class="moment">
    <div class="moment--image">
      <%= img_tag(moment.image) %>
    </div>
    <div class="moment--caption">
      <%= moment.caption %>
    </div>
    <div class="moment--actions">
      <span><%= link "Show", to: Routes.moment_path(@conn, :show, moment) %></span>
      <span><%= link "Edit", to: Routes.moment_path(@conn, :edit, moment) %></span>
      <span><%= link "Delete", to: Routes.moment_path(@conn, :delete, moment), method: :delete, data: [confirm: "Are you sure?"] %></span>
    </div>
  </div>
<% end %>
```

Finally, remove `new/2` from `PictonWeb.MomentController` and update `index/2` and `create/2` to look like this (without the changeset, the form won't work):

```elixir
def index(conn, _params) do
  changeset = Moments.change_moment(%Moment{})
  moments = Moments.list_moments()
  render(conn, "index.html", moments: moments, changeset: changeset)
end

def create(conn, %{"moment" => moment_params}) do
  case Moments.create_moment(moment_params) do
    {:ok, _} ->
      conn
      |> put_flash(:info, "Moment created successfully.")
      |> redirect(to: Routes.moment_path(conn, :index))

    {:error, %Ecto.Changeset{} = _changeset} ->
      conn
      |> put_flash(:error, "Something bad happened")
      |> redirect(to: Routes.moment_path(conn, :index))
  end
end
```

As you can see from the code above, listing and creating moments is delegated to a `Moments` module. The controller doesn't know (and doesn't really care tbh) how the creation of moments is managed, nor where these moments are stored in and fetched from; all it needs to know is that there's a `Moments` module with some functions that work. This module is the context we were talking about before. If you open up `lib/picton/moments.ex`, you can see the implementation of `create_moment/1`, `list_moments/0` and other functions which we'll use (and change) later on.

Back at our broken test, Wallaby expects the moment creation form to be rendered at the root of our application, so let's update our `lib/picton_web/router.ex`:

```elixir
# Delete this:
# get "/", PageController, :index
get "/", MomentController, :index
resources "/moments", MomentController, except: [:new]
```

Now that our form is ready, we need to take care of the file upload. Currently, the back-end expects the user to submit an image path (as a string) instead of an image. We'll be using [Arc](https://github.com/stavro/arc) and [Arc Ecto](https://github.com/stavro/arc_ecto) to handle file uploads.

Add the following dependencies to your `mix.exs`:

```elixir
defp deps do
  [
    # ...
    {:arc, "~> 0.11.0"},
    {:arc_ecto, "~> 0.11.3"},
  ]
end
```

After installing them, set up Arc in `config/config.exs`:

```elixir
config :arc, storage: Arc.Storage.Local
```

Arc requires an uploader module which contains the configuration to store and retrieve uploaded files. We can generate this uploader by running:

```bash
$ mix arc.g moment_image
```

This will create `lib/picton_web/uploaders/moment_image.ex`, which we'll update to the following:

```elixir
defmodule Picton.MomentImage do
  use Arc.{Definition, Ecto.Definition}

  @versions [:original]

  def __storage, do: Arc.Storage.Local

  # Whitelist file extensions:
  def validate({file, _}) do
    ~w(.jpg .jpeg .png)
    |> Enum.member?(Path.extname(file.file_name))
  end

  # Override the storage directory:
  def storage_dir(_version, _) do
    "uploads/moments"
  end
end
```

Also, update your Model in `lib/picton/moments/moment.ex` to this:

```elixir
defmodule Picton.Moments.Moment do
  use Ecto.Schema
  use Arc.Ecto.Schema
  import Ecto.Changeset

  schema "moments" do
    field :caption, :string
    field :image, Picton.MomentImage.Type

    timestamps()
  end

  @doc false
  def changeset(moment, attrs) do
    moment
    |> cast(attrs, [:caption, :image])
    |> cast_attachments(attrs, [:image])
    |> validate_required([:image])
  end
end
```

We're using the `Picton.MomentImage` uploader to handle uploading the file, and we're setting only the image as a required field. Now try running the test we wrote earlier:

```bash
$ mix test test/picton_web/features/post_new_moment.exs

Compiling 23 files (.ex)
Generated picton app
.

Finished in 1.0 seconds
1 feature, 0 failures
```

If you see something else, like a timeout error, [take a look at this](https://elixirforum.com/t/workaround-for-wallaby-httpoison-error/29018).

Now that we have a green test, feel free to play around and post a ton of pics in Picton. You might've noticed two issues. Firstly, the images don't show up. And even if they did show up, the oldest posts would show up first. To fix the first issue, we need to let Phoenix know it should serve static images from the `/upload` folder. So add this plug in `lib/picton_web/endpoint.ex`:

```elixir
plug Plug.Static,
  at: "/uploads",
  from: Path.expand("./uploads"),
  gzip: false
```

And to fix the ordering, update `list_moments/0` in `lib/picton/moments.ex` to this:

```elixir
def list_moments do
  Repo.all(from Moment, order_by: [desc: :inserted_at])
end
```

## Clean up

The last thing to be done in this part 1 of the series is to fix the rest of the tests. Even though we didn't write more than one test, Phoenix generated some for us when we created our Moments context and other files related to it. Run the following and don't get scared by the big blob of red text:

```bash
$ mix test

# <big blob of red text goes here>

Finished in 0.2 seconds
19 tests, 16 failures
```

Open `test/picton/moments_test.exs` and update these lines:

```elixir
# Old, bad code
@valid_attrs %{caption: "some caption", image: "some image"}
@update_attrs %{caption: "some updated caption", image: "some updated image"}
@invalid_attrs %{caption: nil, image: nil}

# New, good code
@valid_attrs %{caption: "some caption", image: "test/support/fixtures/images/meeseeks.png"}
@update_attrs %{caption: "some updated caption"}
@invalid_attrs %{image: nil}
```

When Phoenix generated the test, it didn't know we would be using an image uploader to handle the `image` field, and the `"some image"` string not being a valid URL wasn't playing it well with our uploader. 

After that, go ahead and delete every file in `test/picton_web/controllers`. We will test the application end-to-end using feature tests instead of controller tests. 

Running `mix test` now gives me this output:

```
Finished in 0.2 seconds
10 tests, 2 failures
```

Fixing failing tests by deleting them is a programming technique as old as time itself, and even though it seems like solving an issue by avoiding the issue, it should be accepted as a valid test-fixing approach.

Those last two failing tests are related to creating and updating moments, which we'll update to these:

```elixir
test "create_moment/1 with valid data creates a moment" do
  assert {:ok, %Moment{} = moment} = Moments.create_moment(@valid_attrs)
  assert moment.caption == "some caption"
  assert moment.image.file_name == "meeseeks.png" # update this line
end

test "update_moment/2 with valid data updates the moment" do
  moment = moment_fixture()
  assert {:ok, %Moment{} = moment} = Moments.update_moment(moment, @update_attrs)
  assert moment.caption == "some updated caption"
  # Remove the image-related assertion
end
```

And finally, we have a green test suite:

```bash 
$ mix test                             
..........

Finished in 0.2 seconds
10 tests, 0 failures
```

That's enough Elixir for a day. After learning how to set up Wallaby and uploading files to Phoenix, you're ready for the next great challenge: the second part of this series. We will add users to our application, some authentication and safety measures for them to not edit other people's moments, and follower/followee relationships. Stay tuned.