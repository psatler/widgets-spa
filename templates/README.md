# Endpoints and views 
- GET `/` displays a list of widgets and users at the landing page 
- GET `/users` displays a list of users. Each user has a method of clicking to viewing its details
- GET `/user/:id` displays a single user in detail
- GET `/widgets` displays a list of widgets. Each widget has a method of clicking to viewing its details
- POST `/widgets` creates a new widget
- GET `widgets/:id` displays a single widget in detail. It also has a method for updating its information
- PUT `widgets/:id` Updates an existing widget from form. In fact, this endpoint was built as a POST request in the code, but PUT it is the one suggested to be used
- GET `widgets/:id/edit` Form used for updating a single widget


