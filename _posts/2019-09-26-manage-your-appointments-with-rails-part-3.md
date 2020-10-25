---
layout: post
title: Manage your appointments with Rails, Part 3
date: '2019-09-04'
summary: 
comments: true
category:
  - rails
---

Welcome back to the series. You being here means I’m probably doing something right. If you have any suggestions on what could be done different, or if you think a certain feature would be good to have, or if you feel following the code samples is too solitary and you want to be included more into the series, just drop me a line and let me know.


As I said in the previous [episode of the series](/manage-your-appointments-with-rails-part-2/), we’ll continue with the Slot model. So, let’s.

## Mr. Sandman, bring me a model
We’ve already lined out what the slots table should store. We need to store this information:

 - the scheduled date and time for the slot
 - the host of the event
 - the guest’s email (so we can notify them the event is about to start)
 - notes added from the host for the guest

All this can be done by a single command in your terminal:


{% highlight bash %}
$ rails g model Slot scheduled_at:datetime user:references notes:text guest_email:string
{% endhighlight %}

This command not only creates the model for us, but it also adds a migration file and a spec file for the created model:

{% highlight ruby %}
# app/models/slot.rb
class Slot < ApplicationRecord
  belongs_to :user
end

# db/migrate/XXX_create_slots.rb
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

# spec/models/slot_spec.rb
require 'rails_helper'

RSpec.describe Slot, type: :model do
  pending "add some examples to (or delete) #{__FILE__}"
end
{% endhighlight %}

Let’s proceed by running the migration (`rails db:migrate`) and let’s add some specs for the slots:

{% highlight ruby %}
# spec/factories/slots.rb
FactoryBot.define do
  factory :slot do
    user
    guest_email { FFaker::Internet.email }
    scheduled_at { 10.minutes.from_now }
    notes { '' }
  end
end

# spec/models/slot_spec.rb
require 'rails_helper'

RSpec.describe Slot, type: :model do
  subject { build :slot }

  it { should belong_to(:user).dependent(:destroy).required }

  # email validation
  it { should validate_presence_of(:guest_email) }
  it { should allow_value('test@mail.com').for(:guest_email) }
  it { should_not allow_value('mail.com').for(:guest_email) }
  it { should_not allow_value('test').for(:guest_email) }

  it 'is valid when scheduled at least 5 minutes from now' do
    subject.scheduled_at = 2.minutes.ago
    expect(subject).not_to be_valid

    subject.scheduled_at = 5.minutes.from_now
    expect(subject).to be_valid
  end

  it 'is not valid when scheduled in the past' do
    subject.scheduled_at = 5.minutes.ago
    expect(subject).not_to be_valid
  end
end
{% endhighlight %}

What we’re specifying here is that:

 - When a user is deleted from the DB, all their scheduled slots are also deleted
 - The guest email looks like an email address
 - The slots can only be scheduled to at least 5 minutes later

To make the tests green, here’s what we’ll do to the Slot model:

{% highlight ruby %}
class Slot < ApplicationRecord
  belongs_to :user, dependent: :destroy, required: true

  validates :guest_email, presence: true,
                          format: { with: /\A[A-Z0-9._%a-z\-]+@(?:[A-Z0-9a-z\-]+\.)+[A-Za-z]{2,4}\z/ }

  validate :scheduled_in_the_future

  private

  def scheduled_in_the_future
    seconds_params = { sec: 0 }
    return if scheduled_at.change(seconds_params) >= 5.minutes.from_now.change(seconds_params)

    errors.add(:scheduled_at, 'should be at least 5 minutes from now')
  end
end
{% endhighlight %}

Take a close look at the two different methods we’re using for validation: `validate` and `validates`. The difference between them is that `validates` uses Rails’ built-in validation options, while the other one, `validate`, requires a method (in our case, `scheduled_in_the_future`) that is used for the validation of the object.

And with this code, our specs are green again.

## Give a purpose to the Dashboard
Last time, we created a “dashboard”, but that’s fairly empty. I want to add some functionality to the dashboard, such as listing all the upcoming events (the booked slots), a list of all available free slots (with a way of updating/deleting existing slots) and a way of creating new slots:

![](https://miro.medium.com/max/1400/1*_kdqWvhiEPaXqwosiKr7TA.png)

Using MaterializeCSS, I’m going to build the above view. Try doing the same, [here’s my diff for that](https://github.com/aziflaj/skeduler/commit/c894b5db3257d4879fa342e4017c1439dbbd5e2e). I created partials for each tab, which will make it easier to add Ruby code in the future.

Now, we don’t really have a way to check if one slot is booked by a guest or not. We can assume that a slot with a `guest_email` that is not empty (or `nil`) is a booked slot. Translated into Ruby code, that’d be:

{% highlight ruby %}
class Slot < ApplicationRecord
  # ...
  def booked?
    guest_email.present?
  end
  # ...
end
{% endhighlight %}

In order to manage some more complex queries, e.g. selecting all the user slots that don’t have this guest_email set into the record, we create a Query object.

> Q: Why using Query objects at all?
> A: We want our code as much readable as possible. We can put this code inside the User model and call it as user.free_slots, but if we continue adding similar functionality to the model, then it’ll grow into an unreadable blob of code. Keeping one Query object for each “important” query in the app improves readability, helps during debugging, and complies to the Single Responsibility Principle of [SOLID](https://en.wikipedia.org/wiki/SOLID).

{% highlight ruby %}
# app/queries/free_slots_for_user.rb
module FreeSlotsForUser
  class Query
    def initialize(user)
      @user = user
    end

    def call
      Slot.where(user: user, guest_email: [nil, ''])
    end

    private

    attr_reader :user
  end
end
{% endhighlight %}

Let’s also add some specs for this class:

{% highlight ruby %}
# spec/queries/free_slots_for_user_spec.rb
require 'rails_helper'

RSpec.describe FreeSlotsForUser::Query do
  subject { described_class.new(user) }
  let(:user) { create :user }

  before do
    # Three free slots
    3.times { create :slot, user: user, guest_email: [nil, ''].sample }

    # Three booked slots
    3.times { create :slot, user: user, guest_email: FFaker::Internet.email }
  end

  it 'filters all free slots' do
    expect(subject.call.count).to eq(3)
  end
end
{% endhighlight %}

At this moment I found out that we should remove some specs from the `slot_spec.rb`; we want to allow `nil`/empty values for `guest_email`. So remove the presence validator from the spec file and change the model to look like this:

{% highlight ruby %}
class Slot < ApplicationRecord
  # ...
  validates :guest_email, allow_blank: true,
                          format: { ... } # don't change the format
  # ...
end
{% endhighlight %}

Now, in the dashboard, we want to show these free slots and the upcoming booked slots as well. Add an `UpcomingSlotsForUser::Query` class yourself, and create the following controller:

```ruby
# app/controllers/slots_controller.rb
class SlotsController < ApplicationController
  def free
    slots = FreeSlotsForUser::Query.new(current_user).call

    respond_to do |format|
      format.html { render 'slots/_free', slots: slots }
      format.json { render json: slots }
    end
  end

  def upcoming
    slots = UpcomingSlotsForUser::Query.new(current_user).call

    respond_to do |format|
      format.html { render 'slots/_upcoming', slots: slots }
      format.json { render json: slots }
    end
  end
end

# config/routes.rb
Rails.application.routes.draw do
  # ...
  get 'free_slots', to: 'slots#free'
  get 'upcoming_meetings', to: 'slots#upcoming'
  # ...
end
```


We made our controller to return either HTML content (when `/free_slots` is visited) or JSON content (when `/free_slots.json` is visited).

> If you get an error saying that Rails can’t load the query classes, add this in your `config/application.rb`:
> `config.autoload_paths << Rails.root.join('app', 'queries')`

Now, visiting one of the above configured routes will return the required view.

## HTML is heavy, bandwidth is limited
The title is self-explanatory; if we can return smaller responses when a URL is visited, the user consumes less data from their plan, the lighter responses make the app faster, etc. During the last 5 years, more and more people have been building Single Page Applications (SPA), where the first request sent to the server downloads a blob of JS necessary for running the application, and the rest of the client-server communication is done through JSON. JS frameworks/libraries like Angular, React, Vue, etc., make the frontend management a breeze. But we don’t need any of those to lighten the responses. We already built JSON responses for our `SlotsController`, so let’s consume these endpoints from our frontend.

Instead of returning HTML from our controller, we’ll just return JSON responses:

```ruby
class SlotsController < ApplicationController
  respond_to :json

  def free
    slots = FreeSlotsForUser::Query.new(current_user).call

    render json: slots
  end

  def upcoming
    slots = UpcomingSlotsForUser::Query.new(current_user).call

    render json: slots
  end
end
```

Now, we don’t need to append `.json` to our URLs to get JSON data; we return json by default. To show the fetched data in the browser, we firstly add these ul tags:

```html
<!-- app/views/slots/_upcoming.html.erb -->
<ul class="upcoming-meetings">
</ul>

<!-- app/views/slots/_free.html.erb -->
<ul class="free-slots">
</ul>
```

Then, we make `app/assets/javascripts/dashboard.js` to look like this:

```javascript
document.addEventListener('DOMContentLoaded', function () {
  loadData().then(() => populateViews());

  const tabs = document.querySelectorAll('.tabs');
  M.Tabs.init(tabs);

  const fabs = document.querySelectorAll('.fixed-action-btn');
  M.FloatingActionButton.init(fabs);
});

const loadData = async () => {
  const free_slots_response = await fetch('/free_slots');
  const upcoming_meetings_response = await fetch('/upcoming_meetings');

  window.slots = {
    upcoming: await upcoming_meetings_response.json(),
    free: await free_slots_response.json(),
  };
};

const populateViews = () => {
  const slots = window.slots;

  populateUpcomingMeetings(slots.upcoming);
  populateFreeSlots(slots.free);
};

const populateUpcomingMeetings = (data) => {
  const template = `<li>Meeting with <a href="mailto:$guest_email">$guest_email</a> at $time</li>`;
  const dateOptions = {hour: '2-digit', minute: '2-digit', day: '2-digit', month: 'short', year: 'numeric'};

  document.querySelector('ul.upcoming-meetings').innerHTML = data.map(item => {
    const formattedDate = (new Date(item.scheduled_at)).toLocaleString('en-US', dateOptions)
    return template.replace('$time', formattedDate)
      .replace(/\$guest_email/g, item.guest_email)
  }).join('');
};

const populateFreeSlots = (data) => {
  const template = `<li>Free slot at $time</li>`;
  const dateOptions = {hour: '2-digit', minute: '2-digit', day: '2-digit', month: 'short', year: 'numeric'};

  document.querySelector('ul.free-slots').innerHTML = data.map(item => {
    const formattedDate = (new Date(item.scheduled_at)).toLocaleString('en-US', dateOptions)
    return template.replace('$time', formattedDate)
  }).join('');
};
```

That’s a lot of code, so let’s explain the most important bits:
 - The `loadData()` function asynchronously (i.e. doesn’t block the flow) fetches all the necessary data from the server and assigns the data to the `window.slots` global variable. Consider this variable as a global state for the application, a single source of truth for all the slots data.
 - After the data is loaded, the views are populated using the `populateViews()` function. This function maps the JSON data to a given template, joins all those templates into a single HTML piece, and injects that HTML into our views

There are also other changes in my diff, so [head there and check it all out](https://github.com/aziflaj/skeduler/commit/6d6f21020b0e0cafadbb69a59884875cb2f8d454). My result is this:


![upcoming meetings](https://miro.medium.com/max/1400/1*FH_bxYnkJnS6tqtNgRWkTg.png)

![free slots](https://miro.medium.com/max/1400/1*dag8f3WMgLtxZgZHs6fa8g.png)

## And God said “Let there be…
…slot creation”. And the developer implemented said functionality. And God saw it, and it was asynchronous.

We’ll now give functionality to the FAB that we have in our page. By clicking it, the user (event host) will see a form which allows them to create a free slot. Let’s now create a new partial view `app/views/slots/_form.html.erb` with the following code:

```erbruby
<div id="slot-creation-modal" class="modal">
  <%= form_with model: Slot.new, html: { class: 'new-slot' } do |f| %>
    <div class="modal-content">
      <h4>Create a new slot</h4>
      <div class="row">
        <div class="input-field col s12">
          <i class="material-icons prefix">date_range</i>
          <%= f.text_field :scheduled_at_date, class: 'datepicker', placeholder: 'Scheduled date' %>
        </div>

        <div class="input-field col s12">
          <i class="material-icons prefix">access_time</i>
          <%= f.text_field :scheduled_at_time, class: 'timepicker', placeholder: 'Scheduled time' %>
        </div>
        <div class="input-field col s12">
          <i class="material-icons prefix">event_note</i>
          <%= f.text_area :notes, placeholder: 'Notes', rows: 10 %>
        </div>
      </div>
    </div>
    <div class="modal-footer">
      <%= f.submit 'Create slot', class: 'modal-close waves-effect waves-green btn-flat' %>
    </div>
  <% end %>
</div>
```

Render this partial anywhere you want inside the `app/views/dashboard/index.html.erb` file. To actually create the slot, here’s what we do:

```ruby
app/controllers/slots_controller.rb
class SlotsController < ApplicationController
  protect_from_forgery with: :null_session

  def create
    slot = Slots::Builder.new(current_user, params).call

    if slot.save
      redirect_to dashboard_path
    else
      redirect_to dashboard_path, alert: slot.errors.full_messages.join("\n")
    end
  end
  # ...
end
app/services/slots/builder.rb
module Slots
  class Builder
    def initialize(user, params)
      @user = user
      @params = params
    end

    def call
      Slot.new(slot_params)
    end

    private

    attr_reader :user, :params

    def slot_params
      { user: user,
        scheduled_at: scheduled_at,
        notes: safe_params[:notes] }
    end

    def safe_params
      params.require(:slot).permit(:scheduled_at_date, :scheduled_at_time, :notes)
    end

    def scheduled_at
      DateTime.parse("#{safe_params[:scheduled_at_date]} #{safe_params[:scheduled_at_time]}")
    end
  end
end
```

We’re delegating slot building to a Service Object. Just like the case with Query objects, Services are blocks of code that we use to improve readability and reusability. Usually, service objects are where the business logic goes. Here’s [the whole diff](https://github.com/aziflaj/skeduler/commit/d53be482c15433dd8a2a9bbb9e9f1c29fda82b7e) for what I’ve done.

## To sumarize
You saw how Rails models are created using generators, how to query the DB for data, how to asynchronously fetch data from the server, and how to use service objects to modularise your app.

Next time we’ll do some refactoring on the front-end side. We’ll jump into Webpacker (Rails way of managing Webpack-based frontend modules) and Stimulus.js (a small JS library build by Rails’ team), and we’ll add some more functionality to the app. Stay tuned.

_Previously posted on [my Medium blog](https://medium.com/@aziflaj/manage-your-appointments-with-rails-3-db1313636dfe)_

