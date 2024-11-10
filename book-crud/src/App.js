import React, { useState } from "react";
import BookList from "./components/BookList";
import BookForm from "./components/BookForm";

function App() {
  const [bookToEdit, setBookToEdit] = useState(null);

  const handleEdit = (book) => {
    setBookToEdit(book);
  };

  const handleFormSubmit = () => {
    setBookToEdit(null);
  };

  return (
    <div className="App">
      <h1>Book CRUD App</h1>
      <BookForm bookToEdit={bookToEdit} onFormSubmit={handleFormSubmit} />
      <BookList onEdit={handleEdit} />
    </div>
  );
}

export default App;
