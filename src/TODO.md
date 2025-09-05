# DOING
Refactor how the models access the database (move the Tx creation to the models)
Refactor the Makefiles
Refactor the App start function
Add NIF columns to stores and to users

# TODO
- [ ] On the categories page, add a column of how many times a given category is
used
- [ ] Multiple categories per expense
- [X] Make all the other items update/create page be like expenses
- [ ] Ability to attach a document to an expense
- [ ] Debts table
- [ ] Ability to export data in multiple formats and with a selection of columns
- [ ] Migrate every column named "ExpenseSomething" to "ExpSomething"
- [ ] Make the models just return errors and not do any handling themselves
- [X] Reduce the size of the font file
- [X] Add some migration tool
- [X] Ownership on the image for the expenses webapp dir need to be changed
- [X] The button for deleting share rows isn't working at all unlike the payments
- [X] Use shopspring/decimal for money handling
- [ ] Rename this file
- [ ] Better UX when adding expenses on the payments and share values (default even split)
- [ ] Verification popups for delete actions

### Functionalities
#### Expenses
- Create
- View
- Update
- Delete

#### Categories
- View
- Create
- Delete
- Update

#### Stores
- View
- Create
- Delete
- Update

#### Types
- View
- Create
- Delete
- Update

### Stack
- Go
- golang-migrate
- sqlc
~- sqlite~
- postgresql
- python
- lua

### Patterns
- Service Pattern
