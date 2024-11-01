select * from customers order by age desc limit 1 offset 1; 
 -- to find the second highest aged customer, to find third highest just change offset to 2... and so on for 

SELECT country, COUNT(*) AS customer_count
FROM customers
GROUP BY country; 
-- to find the number of customers in table on the basis of country 
-- output will be like
```
| country | customer_count |
|---------|----------------|
| USA     | 2              |
| UK      | 2              |
| UAE     | 1              |

```

SELECT MAX(age) AS max_age
FROM customers;
-- to find the maximum aged of the customer.