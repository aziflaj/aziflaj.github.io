---

title: Manage your appointments with Rails, Part 1
pubDatetime: 2019-08-28
description: Getting started with Rails by building an Appointment Scheduler
comments: true
icon: icons/rails.png
tags:
  - rails

slug: manage-your-appointments-with-rails-part-1
---

It’s trendy to assume Ruby is dead and the only thing keeping it alive is the Rails framework. It’s also trendy to start every Ruby-related blogpost with some “Ruby is dead” quotes when addressing to non-Ruby devs. Well, this isn’t that kind of blogpost. Instead, consider this a Ruby on Rails primer, a way to get into the Rails train (pun intended). I’m writing this series out of boredom, to give non-Ruby devs an entry point into web development with Rails. I’m not going over what Ruby and Rails are (I won’t even explain MVC to you), neither over how to install them. You can get enough information on that by Googling and if you are that lazy, check [Ruby's](https://www.ruby-lang.org/en/about/) and [Rails'](https://rubyonrails.org/doctrine/) websites. If you need a reference, [take a look at this](https://learnxinyminutes.com/docs/ruby/). You can find the source code for this application [here](https://github.com/aziflaj/skeduler/tree/part1).

## Without losing any time…
…let’s build this Rails application. Speaking of time, I want to build something similar to [Calendly](https://calendly.com/), but with a less beautiful UI (beauty is subjective) and with less features because we don’t have much time.

So, to create a new Rails app, run these in your terminal:

```bash
$ rails new skeduler
$ cd skeduler
$ rails server
```


These simple commands create a new Rails application called Skeduler (no relation to [this other skeduler](https://www.skeduler.com.au/)), then start the server so you can access your application. You can go to [http://localhost:3000/](http://localhost:3000/) and see the greeting Rails page, which shows the running Ruby and Rails version. At this moment, the latest stable versions are respectively 2.6.3 and 6.0.0. I’d suggest you to always use the latest stable version when you start a new application.

## A goal without a plan is just a wish
So let’s plan ahead what we will implement into Skeduler. As is the case with Calendly, we want our users (inside the system, let’s call them hosts) to create some slots: free hours on a day when other people (guests) can book a meeting with them. Without free slots, no booking can happen. After the slot is booked, no other guests can book the same slot (for obvious reasons).

We also want to remind the host and the guests before their scheduled meeting happens, possibly through an email or SMS at some time (e.g. 1 hour) before the meeting.

## Install some gems

Ruby has been around for a while, and the good folks from the community have built more than a few libraries and helpful packages for almost everything you can think of. We call these _gems_, the files that helps us track these dependencies are Gemfile and Gemfile.lock. The first one lists all the dependencies we want in the app, and the other one is used internally from bundler. As you can see, Rails has already added some dependencies for us. Don’t mind them if you’re not already familiar with them. They all are necessary, in a way or another.

```ruby
source 'https://rubygems.org'
git_source(:github) { |repo| "https://github.com/#{repo}.git" }

ruby '2.6.3'

gem 'bootsnap', '>= 1.4.2', require: false
gem 'devise', '~> 4.7' # add this
gem 'jbuilder', '~> 2.7'
gem 'puma', '~> 3.11'
gem 'rails', '~> 6.0.0'
gem 'sass-rails', '~> 5'
gem 'sqlite3', '~> 1.4'
gem 'turbolinks', '~> 5'
gem 'webpacker', '~> 4.0'

group :development, :test do
  gem 'pry-rails', '~> 0.3.9' # add this
  gem 'rspec-rails', '~> 3.8' # add this
end

group :development do
  gem 'web-console', '>= 3.3.0'
  gem 'listen', '>= 3.0.5', '< 3.2'
  gem 'spring'
  gem 'spring-watcher-listen', '~> 2.0.0'
  gem 'rubocop-rails', '~> 2.3', '>= 2.3.1' # add this
end

group :test do
  gem 'capybara', '>= 2.15'
  gem 'selenium-webdriver'
  gem 'webdrivers'
end

# add this
group :production do
  gem 'redis', '~> 4.0'
end

gem 'tzinfo-data', platforms: [:mingw, :mswin, :x64_mingw, :jruby]
```

I changed my Gemfile just a bit; I removed all the unnecessary comments and added some new dependencies. To get them, just run:

```bash
$ bundle install --without production
```

The gems I added are:
 - Devise — We’ll be using this for authenticating users
  - RSpec Rails — We’ll use this to write automated tests for the app
 - Redis (in production) — Needed to run ActionCable, Rails’ way of building WebSocket applications
 - Rubocop — Ruby linter to ensure good coding style
 - Pry — A shell that will become helpful while debugging.

The last step we need to take here is to run:

```bash
$ rm -rf test
$ rails generate rspec:install
$ rails generate devise:install
```

This will remove the default tests and will create a new directory (spec) where all the tests (or specs, if you will) will go. After that, it configures devise for our application. There are some steps listed by devise, which we have to manually follow right now. First, open `config/environments/development.rb` and when you see some lines starting with `config.action_mailer`, add this among those lines:

```ruby
config.action_mailer.raise_delivery_errors = false
config.action_mailer.perform_caching = false
# add this
config.action_mailer.default_url_options = { host: ‘localhost’, port: 3000 }
```

Now, open `app/views/layouts/application.html.erb` and add these two lines just below the <body> tag:

```html
<body>
  <p class="notice"><%= notice %></p>
  <p class="alert"><%= alert %></p>
  <!-- rest of body -->
</body>
```

Finally, run this to generate the devise views in `app/views/devise`:

```bash
$ rails g devise:views # g stands for generate
```

## Your first endpoint
When first arriving at this application, we want to show users a landing page instead of that default Rails welcome page. Let’s start by creating a controller:

```bash
$ rails g controller pages hello --no-controller-specs --no-view-specs
```

This will create a PagesController with a hello method, and a view file with some text on it, and we’re skipping the creation of specs (we can do it by ourself). Now, open `config/routes.rb` and make it look like this:

```ruby
Rails.application.routes.draw do
  root 'pages#hello'
end
```

We use `root` to make sure that whenever the app is opened, the newly-created view will get rendered. Now, start the server and you should see the new view.


## The User model

I expect you to already know what a model is. We’ll let devise handle the creation of the user model, so run:

```bash
$ rails g devise User
$ rails db:migrate
```

This command will
 - Add migration file to create the users table with all the necessary data (the second command runs the migration)
 - Add user model file where device functionality is configured
 - Add the necessary routes for managing user session
 - Add user model spec file

And since we mentioned specs, let’s add some specs for the user.
We already know that the user will have an email, a password with a minimum of 8 characters and a username that can’t be empty or duplicated. So let’s describe this in RSpec talk:

```ruby
require 'rails_helper'

RSpec.describe User, type: :model do
  it 'has a unique email'
  it 'has a password of at least 8 characters'
  it 'has a unique username'
end
```

Now, if you run rspec on your terminal, you’ll see something like:

> 3 examples, 0 failures, 3 pending

Our tests aren’t really implemented (hence the pending status), so let’s implement them naively:

```ruby
require 'rails_helper'

RSpec.describe User, type: :model do
  let(:email) { 'email@example.com' }
  let(:password) { 'password' }

  it 'has a unique email' do
    user1 = User.new(email: email, password: password)
    expect(user1).to be_valid
    user1.save

    user2 = User.new(email: email, password: password)
    expect(user2).not_to be_valid
  end

  it 'has a password of at least 8 characters' do
    user = User.new(email: email, password: 'a' * 7)
    expect(user).not_to be_valid

    user = User.new(email: email, password: 'a' * 8)
    expect(user).to be_valid
  end

  it 'has a unique username' do
    username = 'jamesbond'

    user1 = User.new(email: email, password: password, username: username)
    expect(user1).to be_valid
    user1.save

    user2 = User.new(email: email, password: password, username: username )
    expect(user2).not_to be_valid
  end
end
```

By running rspec you’ll see that some tests pass, and some of them fail for different reasons:

> 3 examples, 2 failures

The first spec passes, since Devise comes with some email validators out of the box. The second spec fails, because Devise validators require passwords to have at least 6 characters. The third spec fails because there’s no `username` field in the `users` table. We can fix the second spec by adding this to the user model:

```ruby
class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable

  # add this
  validates :password, length: { minimum: 8 }
end
```

This validator overrides the one bundled by Devise, so we now accept only passwords with 8 or more characters. As for the third spec, run this

```bash
$ rails g migration add_username_to_users username:string
```

This will create a new migration for us, but we need to change it a bit:

```ruby
class AddUsernameToUsers < ActiveRecord::Migration[6.0]
  def change
    # add the unique: true flag
    add_column :users, :username, :string, unique: true
  end
end
```

Now, run the migration (rails db:migrate) and then run rspec. You’ll now see that all our specs are green.

Finally, let’s make our specs a bit shorter. In the Gemfile, on the test block, add this gem and run bundle install:

```ruby
gem 'shoulda-matchers', '~> 4.1', '>= 4.1.2'
```

Next, create a file called `specs/support/shoulda.rb` and add this code inside (from [the docs](https://github.com/thoughtbot/shoulda-matchers#rspec)):

```ruby
Shoulda::Matchers.configure do |config|
  config.integrate do |with|
    with.test_framework :rspec
    with.library :rails
  end
end
```

Now, open `specs/rails_helper.rb` and uncomment this line of code:

```ruby
Dir[Rails.root.join('spec', 'support', '**', '*.rb')].each { |f| require f }
```

Finally, we can use Shoulda matchers to rewrite the same spec with less code:

```ruby
require 'rails_helper'

RSpec.describe User, type: :model do
  it { should validate_uniqueness_of(:email).case_insensitive }
  it { should validate_length_of(:password).is_at_least(8) }
  it { should validate_uniqueness_of(:username).case_insensitive }
end
```

All those lines of code, condensed into just that. And now, just to verify that we didn’t break anything, run rspec again… and we broke something. The reason this happens is that we only have the uniqueness check for username in the database level, not in the model level. And Shoulda matchers work in the model level. So to fix the failure, add this in the User model:

```ruby
validates :password, length: { minimum: 8 }
# add this
validates :username, uniqueness: { case_sensitive: false }
```

And now, once again, the specs are green.

## To sumarize
In this blogpost you learned to Google what Ruby and Rails are, and you started creating a new Rails application. You installed a bunch of dependencies, some of which you got to use while testing. And you did some Test Driven Development (TDD), which is great. Even though most companies don’t actually do TDD (or BDD) all the time, they implement some sort of automated testing.

Next time we’ll continue with the same codebase and we’ll check how Devise manages user authentication. See you then.

_Previously posted on [my Medium blog](https://medium.com/@aziflaj/manage-your-appointments-with-rails-1-722330e6bc58)_
