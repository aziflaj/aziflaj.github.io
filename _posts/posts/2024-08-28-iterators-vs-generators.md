---
layout: post
title: "Iterators VS Generators: Go's latest YAGNI feature"
date: '2024-08-28'
image:
  path: https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20240828/iterators.jpg?raw=true
  alt: "Iterators are confusing"
---

The new release of Go v1.23 brought us this new feature called "iterators", or 
["rangefuncs"](https://go.dev/wiki/RangefuncExperiment), or
"range-over-funcs". Nobody knows what the real name is but you might have seen them around the interwebz.
Recently I saw it shared on LinkedIn and every time I see it, there's always the same textbook example:

```go
func Backward[T any](s []T) func(func(int, T) bool) {
  return func(yield func(int, T) bool) {
    for i := len(s)-1; i >= 0; i-- {
      if !yield(i, s[i]) {
        return
      }
    }
  }
}
```

And this is how you use this Brain-Backwards Bomboclat:

```go
s := []string{"world", "hello"}
for i, x := range Backward(s) {
  fmt.Println(i, x)
}
```

_Yielding this and yielding that,_ \
_yield a function and who knows what..._

I don't like it.

I don't like the example, because it doesn't feel real-worldy enough. If I wanted to
iterate backwards, I'd write a loop. No functions, no `yield`ing, no fancy stuff. Just a simple loop.
And that was the whole premise of Go, being a _"simple (not easy) language"_.

I don't like the idea of adding a "range-over-func" feature either, because it doesn't solve any problem with the language.
We already had a way to do this: it was called a "Generator Pattern"
and [it was mentioned 12 years ago](https://youtu.be/f6kdp27TYZs?t=866) when Go was still just a baby.
It was a simple way to iterate over a collection
whose size you don't -- _or can't_ -- know beforehand. It was simple, it was idiomatic, and it was Go.

This new feature reminds me of JavaScript's [generator functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/function*),
which also confused me the first time I saw them used in practice. For a second I thought JS had pointers.

> _Consider this post a rant-driven Public Service Announcement, so you at least get a usable,
sane use-case for range-over-funcs. And also a refresher on the Generator Pattern, Go style._

## Better than Backwards Iteration

Forget that Generator pattern exists. Imagine having to iterate over a collection of indefinite length.
You _could_ use a cursor in an infinite loop, and break the loop when the proverbial `cursor.Next()` returns empty/error/whatever.
But for the sake of justifying the existence of this feature, forget cursors are a thing.
We are not here to learn from the past, we are here to embrace the future,
and to enshittify a good language simply to accommodate developers coming from other languages.

Let's say you're pulling some paginated data from a web service, as some sort of infinite scroll.
There's always a "next page" (until there isn't), and you want to iterate over all pages (until _somehow_ prompted to stop)
and render all the results you get. With range-over-funcs it would look like this:

```go
func FetchPaginatedData[T any](rawURL, pageParam string) func(func(T, error) bool) {
  var page int

  // Make sure the url is a URL
  url, err := url.ParseRequestURI(rawURL)
  if err != nil {
    panic(err) // this ain't gonna fly, bucko!
  }

  // default the page count to 0
  query := url.Query()
  if query.Has(pageParam) {
    page, err = strconv.Atoi(query.Get(pageParam))
    if err != nil { // to err is human, to rectify divine
      page = 0
    }
  }

  // fetch each page and yield the data or the error
  return func(yield func(T, error) bool) {
    var data T
    client := &http.Client{}

    for {
      query.Set(pageParam, strconv.Itoa(page))
      url.RawQuery = query.Encode()

      res, err := client.Get(url.String())
      if err != nil {
        yield(data, err)
        break
      }
      defer res.Body.Close()

      if res.StatusCode != http.StatusOK {
        yield(data, fmt.Errorf("Unexpected status code: %v", res.StatusCode))
        break
      }

      err = json.NewDecoder(res.Body).Decode(&data)
      if err != nil {
        yield(data, err)
        break
      }

      yield(data, nil)

      // Continue to the next page
      page += 1
    }
  }
}
```

This function allows you to `range` over some paginated data, and in each iteration
you will get a different page of data. If something goes wrong, e.g. if your
URL is badly formatted, or the server returns some error code, you will get that
error and the iteration will stop. And for the sake of practicity, we make this
iterator a generic one, so you can fetch and paginate over any type of data.
Here's how you use it:

```go
type MyResp struct {
  // TODO: Implement
}

// You can start paginating from any page
for data, err := range FetchPaginatedData[MyResp]("https://example.com?page=42", "page") {
  if err != nil {
    fmt.Printf("Shit went sideways: %v", err)
    // No need to break
    // Iteration stops in case of error
  }

  // TODO: process data
  fmt.Println(data)
}
```

It is easy to use, sure. If someone else wrote the `FetchPaginatedData` function,
this code is fairly readable. But I find those `yield`s a bit unreadable and confusing.
This implementation only reminds me that we already had a way to do this even without
yielding.

## Old Habits Die Hard

So how did we do this in pre-1.23 Go? How do we refactor this new yielder into a simpler, idiomatic Go generator?

The bulk of the work is going to remain the same.
We still need to do the same URL parsing and sending a `page + 1` request after every "iteration".
But instead of `yield`ing, we will push the fetched data into a channel
and instead of `range`-ing over a function, we will read from the channel until it's closed:

```go
func PaginateOverChannels[T any](rawURL, pageParam string) (<-chan T, <-chan error) {
  dataCh := make(chan T)
  errCh := make(chan error)

  var page int
  url, err := url.ParseRequestURI(rawURL)
  if err != nil {
    panic(err) // this ain't gonna fly, bucko!
  }

  query := url.Query()
  if query.Has(pageParam) {
    page, err = strconv.Atoi(query.Get(pageParam))
    if err != nil { // to err is human, to rectify divine
      page = 0
    }
  }

  go func() {
    var data T
    client := &http.Client{}
    for {
      query.Set(pageParam, strconv.Itoa(page))
      url.RawQuery = query.Encode()

      res, err := client.Get(url.String())
      if err != nil {
        errCh <- err
        break
      }
      defer res.Body.Close()

      if res.StatusCode != http.StatusOK {
        errCh <- fmt.Errorf("Unexpected status code: %v", res.StatusCode)
        break
      }

      err = json.NewDecoder(res.Body).Decode(&data)
      if err != nil {
        errCh <- err
        break
      }

      dataCh <- data

      page += 1
    }

    close(dataCh)
    close(errCh)
  }()

  return dataCh, errCh
}
```

The code is almost the same, with the addition of two channels: one for the data and one for the errors.
Also, I took the liberty of pulling the data in a goroutine, because _concurrency_. And to use it:

```go
dataChan, errChan := PaginateOverChannels[MyResp]("https://example.com?page=42", "page")

for {
  select {
    case data, ok := <-dataChan:
      if !ok {
        return // or break the outer `for` loop
      }
      fmt.Println(data)

    case err, ok := <-errChan:
      if !ok {
        return
      }
      fmt.Printf("Shit went sideways: %v\n", err)
  }
}
```

Maybe writing a `for` + `select` is not as simple and easy as writing a `range`,
but channels are such a good and core feature of Go and so baked into the brain of whoever
uses the language, it makes no sense (or need, for that being) to add a different feature.

## What are you hiding from us?

Too many things!

I know I left some things out for simplicity and to focus on
the practical use case. One thing is the `iter` package. You see how our 
`FetchPaginatedData` function returns a `func(func(T, error) bool)`? There's a 
new way of doing that. Instead, you should return an `iter.Seq2[T, error]`.
There's also an `iter.Seq[T]` if you need to iterate a single value.

Also, I know the channel approach needs a bit of tweaking to make it stoppable on demand,
but that goes a bit beyond the scope of this example; not the right _context_, if you will.

Anywho, you can read some [more detailed insights about iterators from John Arundel here](https://bitfieldconsulting.com/posts/iterators)
and you might also want to see [how range-over-func is used to implement a cursor-y iterator for a CouchDB library](https://boldlygo.tech/posts/2024-07-18-range-over-func/).
