# Problem Statement: Customer Order Management Microservice
A company requires a Customer Order Management Microservice that allows internal applications to handle customer orders efficiently. The service will provide a REST API for managing customer profiles, placing orders, viewing order status, and performing administrative tasks.

### Requirements

1. Customer Management:
    - Create Customer: Accept new customer registrations with information like name, email, phone number, and address.
    - Get Customer Details: Retrieve details of a specific customer by their unique ID.
    - Update Customer Information: Allow customers to update their profile information.
    - Delete Customer Profile: Soft delete a customer profile, marking it as inactive.
2. Order Management:
    - Place Order: Allow customers to place an order by providing product details (e.g., product ID, quantity).
    - Get Order Status: Retrieve the status of a specific order by its order ID (e.g., Pending, Shipped, Delivered).
    - Update Order: Update the order quantity or product details before it is processed.
    - Cancel Order: Allow customers to cancel an order if it’s not yet shipped.
3. Administrative Functions:
    - List All Orders: Provide a paginated list of all orders placed, with filters for order status and customer ID.
    - Get Customer Orders: Retrieve all orders for a specific customer, along with their status.

### Technical Requirements

- Framework: Use the Go-Chi framework for building the REST API.
- Data Persistence: Use a database (e.g., PostgreSQL) to store customer and order data.
- Authentication: Secure certain endpoints (like order placement and update) using token-based authentication (e.g., JWT).
- Validation: Validate incoming request data for all endpoints.
- Error Handling: Implement consistent error handling and return meaningful error messages.
- Logging: Integrate structured logging for request and response data, including error logs.
- Testing: Write unit tests and integration tests for core functionalities.
- Documentation: Use Swagger or similar tools to document API endpoints for developers.

### Performance and Scalability

The microservice should be designed to handle high request rates and should be easily scalable to accommodate future growth in order volume.