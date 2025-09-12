# WIP
This is a simple webapp I use to track expenses. It has many holes and things
that don't work properly, but the current feature set is:
- Users (Authentication)
- Multiple APIs: Lua API, REST API that returns JSON and of course the core
`expenses` package itself that can be used separately
- Insert expenses via QR code (at least the data that is available, am working
on making the experience as seamless as possible)
- Share expenses arbitrarily with others
- Check the balance on your own expenses


There are many features to be implemented!


# BUGS
- Inserting an expense where the owner has no payments results in a panic


# DOING


# TODO
- [X] Users/Authentication
- [ ] Authorization
- [X] API (not complete, but going)
- [ ] Add more logging (and configure the logger)
- [ ] Need to add many UTs
- [ ] Need to think about where the authentication endpoints are its separation
of concerns
- [ ] Create an interface to get data from a QR code so that we can have both
a web API and our go package
- [ ] If unable to connect to the database, try again, after X seconds
- [ ] On the categories page, add a column of how many times a given category is
used
- [ ] Multiple categories per expense
- [ ] Associate categories to a given store (automate giving a category when inserting via QR)
- [X] Make all the other items update/create page be like expenses
- [ ] Ability to attach a document to an expense
~~- [ ] Debts table~~
- [ ] Ability to export data in multiple formats and with a selection of columns
- [ ] Migrate every column named "ExpenseSomething" to "ExpSomething"
- [X] Make the models just return errors and not do any handling themselves
- [X] Reduce the size of the font file
- [X] Add some migration tool
- [X] Ownership on the image for the expenses webapp dir need to be changed
- [X] The button for deleting share rows isn't working at all unlike the payments
- [X] Use shopspring/decimal for money handling
- [X] Rename this file
- [ ] Better UX when adding expenses on the payments and share values (default even split)
- [ ] Verification popups for delete actions
- [ ] Allow user X to insert expense of user Y - make user Y have to approve it


### Stack
- Go
- Gin Web Framework
- golang-migrate
- sqlc
~- sqlite~
- postgresql
- python
- lua


### Patterns
- Service Pattern
- Repository Pattern


### Architecture
- MVC


## Thanks to
This excellent write up:
- [Database Transactions in GO](https://threedots.tech/post/database-transactions-in-go/),
by [m110](https://github.com/m110)
