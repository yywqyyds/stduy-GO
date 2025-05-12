import React, { useState, useEffect } from 'react';
import TodoItem from './components/TodoItem';
import { TodoContext } from './components/TodoContext';
import type { Todo } from './components/TodoContext';
import './styles.css';

const App: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [input, setInput] = useState('');
  const [darkMode, setDarkMode] = useState(false);

  useEffect(() => {
    const saved = localStorage.getItem('todos');
    if (saved) setTodos(JSON.parse(saved));
  }, []);

  useEffect(() => {
    localStorage.setItem('todos', JSON.stringify(todos));
  }, [todos]);

  const addTodo = () => {
    if (!input.trim()) return;
    const newTodo: Todo = {
      id: Date.now(),
      text: input.trim(),
      time: new Date().toLocaleString(),
      completed: false,
    };
    setTodos([newTodo, ...todos]);
    setInput('');
  };

  const toggleComplete = (id: number) => {
    setTodos(todos.map(todo => todo.id === id ? { ...todo, completed: !todo.completed } : todo));
  };

  const deleteTodo = (id: number) => {
    setTodos(todos.filter(todo => todo.id !== id));
  };

  const toggleTheme = () => setDarkMode(prev => !prev);

  return (
    <TodoContext.Provider value={{ todos, toggleComplete, deleteTodo }}>
      <div className={`app ${darkMode ? 'dark' : ''}`}>
        <h1 className="title">Todo List</h1>
        <button className="theme-toggle" onClick={toggleTheme}>
          {darkMode ? '🌞 Light' : '🌙 Dark'}
        </button>
        <input
          className="input"
          placeholder="What needs to be done?"
          value={input}
          onChange={e => setInput(e.target.value)}
          onKeyDown={e => e.key === 'Enter' && addTodo()}
        />
        <ul className="todo-list">
          {todos.map(todo => (
            <TodoItem key={todo.id} todo={todo} />
          ))}
        </ul>
      </div>
    </TodoContext.Provider>
  );
};

export default App;
