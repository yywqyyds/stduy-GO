import React from 'react';
import { useTodoContext } from './TodoContext';
import type { Todo } from './TodoContext';
import './TodoItem.css';

type Props = {
  todo: Todo;
};

const TodoItem: React.FC<Props> = ({ todo }) => {
  const { toggleComplete, deleteTodo } = useTodoContext();

  return (
    <li className={`todo-item ${todo.completed ? 'completed' : ''}`}>
      <input
        type="checkbox"
        checked={todo.completed}
        onChange={() => toggleComplete(todo.id)}
      />
      <span className="text">{todo.text}</span>
      <span className="time">{todo.time}</span>
      <button className="delete" onClick={() => deleteTodo(todo.id)}>🗑️</button>
    </li>
  );
};

export default TodoItem;
