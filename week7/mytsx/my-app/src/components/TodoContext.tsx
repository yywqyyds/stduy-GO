import  { createContext, useContext } from 'react';

export type Todo = {
  id: number;
  text: string;
  time: string;
  completed: boolean;
};

type TodoContextType = {
  todos: Todo[];
  toggleComplete: (id: number) => void;
  deleteTodo: (id: number) => void;
};

export const TodoContext = createContext<TodoContextType | undefined>(undefined);

export const useTodoContext = () => {
  const context = useContext(TodoContext);
  if (!context) throw new Error('useTodoContext must be used within TodoProvider');
  return context;
};
