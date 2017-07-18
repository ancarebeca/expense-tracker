# Expense-tracker

Given a csv with the last summary of a bank account, it organises this expenses in categories 
(bills, shopping, eating out …) and then it will be saved into an Expense database.  
After that, these data will be represented in a dashboard

Services:

- Extraction: Load en memory a csv file from a bank account
- Transformation: Normalized csv data
- Categorize: Set categories to each expense (In progress)
- Load: Load data into a database
- Report: Show the information in a dashboard (In progress)

## Installation:

1. Install [golang](https://golang.org/doc/install)
 
2. Dependencies:

```
    make install
```

## Configuration:
1. Define configuration file:

```
#expense-tracker/config
    user_db: "your_db"
    pass_db: "your_db_password"
    database: "your_db"
    file_path: "path_your_bank_statement"
```

2. Link your description keyword with a category in  `expense-tracker/config/categories.yaml`: 

You have to link some relevant word from your bank statements with one category.
For example if in your bank statement in the description field you have 'Tfl.gov.uk/cp' you 
could categorise 'tfl' as transport. 

```
tfl: transport
my groceries shop: groceries
my favourite restaurant: eating out
atm machine: cash
thames water: bills
```

## Run tests: 

```
    make install
```

## Run app: 

```
    make all
```
