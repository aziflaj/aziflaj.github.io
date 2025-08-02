---

title:      "Week #11: Jump Starting Sinatra"
pubDatetime:       2016-03-04
comments:   true
description:    I developed a simple Sinatra application, after reading through "Jump Start Sinatra" by Darren Jones (apparently, there's a footballer with the same name; he didn't write a book). While I've been playing before with Ruby, using Rails and Chef, this is my first time with Sinatra.
tags:
    - project52
    - ruby
    - sinatra

slug: week-11-jump-starting-sinatra
---

No more exams! As I wrote in the end of my last blogpost, I had to take one month off from the challenge because of the finals. Now I'm free to go on with the Project 52, to learn more and become better. This week, I developed a simple Sinatra application, after reading through "Jump Start Sinatra" by [Darren Jones](https://twitter.com/daz4126) (apparently, there's [a footballer](https://en.wikipedia.org/wiki/Darren_Jones) with the same name; he didn't write any book).

If you're not familiar with it, [Sinatra](http://www.sinatrarb.com/) is a [Domain-Specific Language (DSL)](https://en.wikipedia.org/wiki/Domain-specific_language) written in Ruby that allows you to create simple web applications quickly and without too much effort. While I've played before with Ruby, using Rails and [Chef](http://aziflaj.github.io/week-3-challenge-cooking-virtual-machines-with-chef/), this is my first time using Sinatra.

## Jump-Start Sinatra, the book
[The book](http://www.sitepoint.com/store/jump-start-sinatra/) is not that big, more or less 150 pages for the reader including code examples as the author takes you through each stage  of developing an entire application through the book. The language is simple and even if you have no previous experience with Ruby, by reading it you can understand a lot about how Sinatra works and how to use it to develop web applications.

After following the instructions in the book, by the end you will have built a CRUD application using SQLite3 in development and PostgreSQL in production and will have it deployed in [Heroku](http://heroku.com/), which implies basic [Git](http://git-scm.com/) skills. You will learn how to use [Sass](http://sass-lang.com/) and [Coffeescript](http://coffeescript.org/) and compile `.scss` and `.coffee` files into `.css` and `.js`. You will learn to use [Slim](http://slim-lang.com/) templating instead of plain erb views. If you want to learn to use Sinatra, this is a must-read book!

## The Gems DB
The application I developed is called **The Gems DB**, which you can find on Github ([aziflaj/thegemsdb](https://github.com/aziflaj/thegemsdb)) and also [deployed on Heroku](http://the-gemstone-database.herokuapp.com/). The application is quite simple: it allows you add new gemstones in the database, check existing ones, update their description and also delete them. I didn't use everything from the book mentioned above actually, as I decided I didn't need some of the components. I used Bootstrap for styling and only one line of CSS, so using Sass was a bit overkill. Also, I didn't use Javascript at all, so using Coffeescript was completely obsolete.

The application uses SQLite3 in development and Postgres in production. It uses [DataMapper](http://datamapper.org/) as ORM, which is fairly simple to use. To create the model of the gemstone stored in the DB, I created a class like this one:

```ruby
require 'dm-core'
require 'dm-migrations'

class Gemstone
  include DataMapper::Resource
  property :id, Serial            # Serial = AutoIncrement
  property :name, String
  property :description, Text
end

# Before using the Gemstone class, we make sure to call this method below
# to be able to use the class
DataMapper.finalize
```

Then, the model is used to retrieve objects from the DB like:

```ruby
require './gemstone'

# ...

get '/:id' do
  @gemstone = Gemstone.get(params[:id])
  if @gemstone
    erb :"gemstones/show", layout: :layout
  else
    not_found
  end
end
```

The above is the block executed when the `/gemstones/:id` route is accessed. If the `@gemstone` variable is not null, the application renders a view (`gemstones/show.erb`) with the information of the gemstone. Otherwise, a block called `not_found` is executed which renders a 404 page. The view mentioned is:

```ruby
<% @title = "#{@gemstone.name } | TGDB" %>

<h2><%= @gemstone.name %></h2>
<a href="<%= "/gemstones/#{@gemstone.id}/edit" %>" class="btn btn-default btn-sm">Edit</a>
<p><%= @gemstone.description %></p>
```

As you can see, the `@gemstone` variable is also passed to the view.

## The environment
Recently, I had to reinstall my operating system (Fedora, just like the one Linus Torvalds uses) because of some issue I had, so all my previous configurations were lost. Now that I have a fresh installation, I decided to not make my OS "dirty" with all these configurations and development tools (except Android). Instead, I am now using a Vagrant machine which is somehow generic and can be used for PHP, Ruby (Rails and Sinatra tested), Node.js and Python (Flask tested) development. The VM has a couple of database management systems installed including PostgreSQL, SQLite3, MySQL, MongoDB and Redis (because you never know). Nginx is also installed, and if you want to use the same VM, you can find [the config file here](https://gist.github.com/aziflaj/b9ba5893f41e58023c6b). All you have to do is drag and drop that config file to [PuPHPet.com](https://puphpet.com/), click the **Create Archive** button on your left, donwload the archive, extract it, and execute this in that directory:

```bash
$ vagrant up && vagrant ssh
```

Some of the ports are forwarded (host => guest):

- 3000 => 3000
- 5000 => 5000
- 8000 => 80
- 33060 => 3306

And that's it for this week! It feels good to come back at this project, and I have some good things in mind for the future. I hope you like the application I developed, and if you have anything to add to it please fork it and send a pull request, or leave any comment below in the comment section.
