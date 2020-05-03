# BestPrice Back-end Test


Requirements
- Use Golang for application logic - **Done**  
- Use MySQL for data storage - **Done**  
- Requests and responses should be in JSON -**Done**  
- All "list" endpoints must support:  
    - Pagination using either page/per_page or offset/limit - **Done**  
    - Sorting all values in ascending or descending order - **Done**  
- Your SQL queries should be performant and prevent SQL injection - **Done**  
- Provide a MySQL DDL so it can be tested locally - **Done**  


Optional  
- Use caching to improve performance of "list" and "read" endpoints, using Memcached or Redis - **Not Done _even though I wanted_**  
- Enforce authentication for "create", "update", and "delete" endpoints - **Done _with JWT_**  
- Unit testing - **Started TDD but stopped due to lack of time _(tests don't run correctly)_**  

There is a SQL file and a postman colection in the docs directory