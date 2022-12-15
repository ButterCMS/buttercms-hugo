# ButterCMS Hugo Starter Project

This is a starter project for Hugo with integration of ButterCMS. This includes a posts generator written in Go which uses the ButterCMS SDK for Go to retrieve posts from ButterCMS.

## Important Notice
This project was created as an example use case of ButterCMS in conjunction with Hugo, and will not be actively maintained. 

If you’re interested in exploring the best, most up-to-date way to integrate Butter into Hugo, you can check out the following resources:

### Starter Projects

The following turn-key starters are fully integrated with dynamic sample content from your ButterCMS account, including main menu, pages, blog posts, categories, and tags, all with a beautiful, custom theme with already-implemented search functionality. All of the included sample content is automatically created in your account dashboard when you sign up for a free trial of ButterCMS.
- [Hugo Starter](https://buttercms.com/starters/hugo-starter-project/)
- [Angular Starter](https://buttercms.com/starters/angular-starter-project/)
- [React Starter](https://buttercms.com/starters/react-starter-project/)
- [Vue.js Starter](https://buttercms.com/starters/vuejs-starter-project/)
- Or see a list of all our [currently-maintained starters](https://buttercms.com/starters/). (Over a dozen and counting!)

### Other Resources
- Check out the [official ButterCMS Docs](https://buttercms.com/docs/)
- Check out the [official ButterCMS API docs](https://buttercms.com/docs/api/)

## Installation

1. Install Hugo.

   You can check on how to install Hugo on your operating system [here](https://gohugo.io/getting-started/installing/)

2. Install Go

   You can check on how to install Go compiler [here](https://golang.org/doc/install)

## Setup

You should replace YOUR_AUTH_TOKEN in go/generate.go with your auth token provided by ButterCMS.


## Usage

1. Run the generate script to retrieve all of the posts from ButterCMS
```bash
go run go/generate.go
```

2. Check the posts if they are loaded into HUGO by running
```bash
hugo serve
```

Web Server will be available at http://localhost:1313/
