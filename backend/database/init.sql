CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0)
);

INSERT INTO
    books (title, author, quantity)
VALUES
    (
        'The Hitchhiker''s Guide to the Galaxy',
        'Douglas Adams',
        42
    ),
    ('1984', 'George Orwell', 100),
    ('The Lord of the Rings', 'J.R.R. Tolkien', 50),
    ('The Catcher in the Rye', 'J.D. Salinger', 34),
    ('To Kill a Mockingbird', 'Harper Lee', 60);