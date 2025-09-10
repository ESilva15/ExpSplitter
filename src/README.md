# WIP
This is a simple webapp I use to track expenses. It has many holes and things
that don't work properly. I use this README internally to keep in mind what
I'm working on.


# BUGS
Inserting an expense where the owner has no payments results in a panic


# DOING
- Auth


# TODO
- [ ] Users/Authentication/Autherization
- [ ] API
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
https://threedots.tech/post/database-transactions-in-go/
https://github.com/m110
