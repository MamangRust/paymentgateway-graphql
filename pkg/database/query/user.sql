-- GetUsersWithPagination: Search Users with Pagination and Total Count
-- name: GetUsersWithPagination :many
-- Purpose: Retrieve users with pagination and total count of users matching the search criteria
-- Parameters:
--   $1: search_term - A search term to filter users by firstname, lastname, or email
--   $2: limit - The maximum number of users to return per page
--   $3: offset - The number of users to skip (for pagination)
-- Returns:
--   - User records matching the search term, including firstname, lastname, and email
--   - A total count of matching users, including all pages (using COUNT(*) OVER())
-- Business Logic:
--   - Filters users by search_term (if provided), allowing case-insensitive search.
--   - Returns paginated results, ordered by `created_at` in descending order.
--   - The total count includes the entire dataset, not limited by pagination.
-- name: GetUsersWithPagination :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;



-- GetActiveUsersWithPagination: Get Active Users with Pagination and Total Count
-- name: GetActiveUsersWithPagination :many
-- Purpose: Retrieve active (non-deleted) users with pagination and total count
-- Parameters:
--   $1: search_term - A search term to filter active users by firstname, lastname, or email
--   $2: limit - The maximum number of active users to return per page
--   $3: offset - The number of active users to skip (for pagination)
-- Returns:
--   - Active user records matching the search term, including firstname, lastname, and email
--   - A total count of active users, including all pages (using COUNT(*) OVER())
-- Business Logic:
--   - Filters users where `deleted_at` is NULL (only active users).
--   - Allows filtering by search term across firstname, lastname, or email.
--   - Returns paginated active users, ordered by `created_at` in descending order.
--   - The total count of active users is calculated, including those that are not currently on the current page.
-- name: GetActiveUsersWithPagination :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- GetTrashedUsersWithPagination: Get Trashed Users with Pagination and Total Count
-- name: GetTrashedUsersWithPagination :many
-- Purpose: Retrieve trashed (soft-deleted) users with pagination and total count
-- Parameters:
--   $1: search_term - A search term to filter trashed users by firstname, lastname, or email
--   $2: limit - The maximum number of trashed users to return per page
--   $3: offset - The number of trashed users to skip (for pagination)
-- Returns:
--   - Trashed user records matching the search term, including firstname, lastname, and email
--   - A total count of trashed users, including all pages (using COUNT(*) OVER())
-- Business Logic:
--   - Filters users where `deleted_at` is NOT NULL (only trashed users).
--   - Allows filtering by search term across firstname, lastname, or email.
--   - Returns paginated trashed users, ordered by `created_at` in descending order.
--   - The total count of trashed users is calculated, including those that are not currently on the current page.
-- name: GetTrashedUsersWithPagination :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NOT NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;



-- GetUserByID: Retrieve a user by their ID
-- name: GetUserByID :one
-- Purpose: Fetch details of a specific user by their unique user_id.
-- Parameters:
--   $1: user_id - The ID of the user to fetch.
-- Returns:
--   - User record matching the user_id with the `deleted_at` column being NULL (indicating the user is active).
-- Business Logic:
--   - Filters the users table to find a user based on their `user_id`.
--   - Ensures the user is active by checking that `deleted_at` is NULL.
-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1 AND deleted_at IS NULL;


-- GetUserByEmail: Retrieve a user by their email
-- name: GetUserByEmail :one
-- Purpose: Fetch a specific user based on their email.
-- Parameters:
--   $1: email - The email of the user to fetch.
-- Returns:
--   - User record matching the provided email with `deleted_at` being NULL (active user).
-- Business Logic:
--   - Filters the users table by email to find a user.
--   - Ensures that the `deleted_at` field is NULL, so only active users are returned.
-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL;

-- SearchUsersByEmail: Search users by email with case-insensitive matching
-- name: SearchUsersByEmail :many
-- Purpose: Allows searching for users whose email matches a given search term (case-insensitive).
-- Parameters:
--   $1: email_search_term - A partial or full email address to search for.
-- Returns:
--   - List of users whose emails match the search term.
--   - The results are ordered by the `created_at` column in descending order.
-- Business Logic:
--   - Uses `ILIKE` to perform a case-insensitive search on the `email` column.
--   - Only returns active users (`deleted_at IS NULL`).
-- name: SearchUsersByEmail :many
SELECT *
FROM users
WHERE
    deleted_at IS NULL
    AND email ILIKE '%' || $1 || '%'
ORDER BY created_at DESC;

-- GetTrashedUserByID: Retrieve trashed user by their ID
-- name: GetTrashedUserByID :one
-- Purpose: Fetch a trashed (soft-deleted) user based on their user_id.
-- Parameters:
--   $1: user_id - The ID of the trashed user to fetch.
-- Returns:
--   - User record matching the user_id where `deleted_at` is not NULL (indicating the user is trashed).
-- Business Logic:
--   - Filters the users table to find a trashed user based on their `user_id`.
--   - Checks that `deleted_at` is NOT NULL to ensure the user is trashed.
-- name: GetTrashedUserByID :one
SELECT *
FROM users
WHERE
    user_id = $1
    AND deleted_at IS NOT NULL;



-- CreateUser: Insert a new user into the users table
-- name: CreateUser :one
-- Purpose: Add a new user to the system.
-- Parameters:
--   $1: firstname - The first name of the user.
--   $2: lastname - The last name of the user.
--   $3: email - The email address of the user.
--   $4: password - The password of the user (hashed).
-- Returns:
--   - The newly created user record.
-- Business Logic:
--   - Inserts a new user record into the `users` table with the current timestamp for `created_at` and `updated_at`.
-- name: CreateUser :one
INSERT INTO
    users (
        firstname,
        lastname,
        email,
        password,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        current_timestamp,
        current_timestamp
    ) RETURNING *;



-- UpdateUser: Modifies user account information
-- Purpose: Update user profile details
-- Parameters:
--   $1: user_id - ID of user to update
--   $2: firstname - Updated first name
--   $3: lastname - Updated last name
--   $4: email - Updated email address
--   $5: password - New hashed password (optional)
-- Returns: Updated user record
-- Business Logic:
--   - Auto-updates updated_at timestamp
--   - Only modifies active (non-deleted) users
--   - Validates email uniqueness
--   - Password field optional (can maintain existing)
-- name: UpdateUser :one
UPDATE users
SET
    firstname = $2,
    lastname = $3,
    email = $4,
    password = $5,
    updated_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL
    RETURNING *;



-- TrashUser: Soft-deletes a user account
-- Purpose: Deactivate user without permanent deletion
-- Parameters:
--   $1: user_id - ID of user to deactivate
-- Returns: The soft-deleted user record
-- Business Logic:
--   - Sets deleted_at timestamp to current time
--   - Only processes currently active users
--   - Preserves all user data for potential restoration
--   - Prevents login while deleted
-- name: TrashUser :one
UPDATE users
SET
    deleted_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL
    RETURNING *;

-- RestoreUser: Recovers a soft-deleted user
-- Purpose: Reactivate a previously deactivated user
-- Parameters:
--   $1: user_id - ID of user to restore
-- Returns: The restored user record
-- Business Logic:
--   - Nullifies the deleted_at field
--   - Only works on previously deleted users
--   - Restores full account access
--   - Maintains all original user data
-- name: RestoreUser :one
UPDATE users
SET
    deleted_at = NULL
WHERE
    user_id = $1
    AND deleted_at IS NOT NULL
    RETURNING *;


-- DeleteUserPermanently: Permanently delete a trashed user from the system
-- name: DeleteUserPermanently :exec
-- Purpose: Permanently delete a trashed user record.
-- Parameters:
--   $1: user_id - The ID of the trashed user to delete permanently.
-- Business Logic:
--   - Deletes the user record from the `users` table permanently.
--   - Only deletes users who have been trashed (`deleted_at IS NOT NULL`).
-- name: DeleteUserPermanently :exec
DELETE FROM users WHERE user_id = $1 AND deleted_at IS NOT NULL;


-- RestoreAllUsers: Restore all trashed users
-- name: RestoreAllUsers :exec
-- Purpose: Restore all soft-deleted users by clearing their `deleted_at` field.
-- Business Logic:
--   - Clears the `deleted_at` field for all trashed users, effectively restoring them.
-- name: RestoreAllUsers :exec
UPDATE users
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;


-- DeleteAllPermanentUsers: Permanently delete all trashed users
-- name: DeleteAllPermanentUsers :exec
-- Purpose: Permanently delete all trashed user records from the database.
-- Business Logic:
--   - Deletes all users who have been trashed (soft-deleted), i.e., where `deleted_at` is not NULL.
-- name: DeleteAllPermanentUsers :exec
DELETE FROM users
WHERE
    deleted_at IS NOT NULL;

