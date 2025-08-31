# DOING
Adding the field `Calculated` to `Share` to avoid computing what each user owes
everytime due to the fractional cents problem
Make a script for the current migration to compute the Calculated value with
our rules

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

### Stack
- Go
- golang-migrate
- sqlc
- sqlite

### Patterns
- Service Pattern
