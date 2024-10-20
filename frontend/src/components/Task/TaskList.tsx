import React, { useEffect, useState } from 'react';
import { Task, TaskUpdate } from '../../types/task';
import { TaskService } from '../../services/taskService';
import TaskForm from './TaskForm';

const TaskList: React.FC = () => {
    const [tasks, setTasks] = useState<Task[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [editingTask, setEditingTask] = useState<Task | null>(null);

    const fetchTasks = async () => {
        try {
            const data = await TaskService.listTasks();
            console.log('Fetched tasks:', data);  // Debugging the task data
            setTasks(data);
        } catch (err) {
            setError('Failed to load tasks');
            console.error('Error fetching tasks:', err);
        }
    };

    const handleDelete = async (ID: number) => {
        console.log(`Deleting task with ID: ${ID}`);  // Debug: Log the ID
        try {
            await TaskService.deleteTask(ID);
            setTasks((prevTasks) => prevTasks.filter((task) => task.ID !== ID));
        } catch (error) {
            console.error('Failed to delete task:', error);
        }
    };

    const handleEdit = (task: Task) => {
        console.log('Editing task:', task.ID);  // Debugging the task to be edited
        setEditingTask(task);
    };

    const handleUpdate = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!editingTask) return;

        try {
            const updatedTask: TaskUpdate = {
                title: editingTask.title,
                description: editingTask.description,
                priority: editingTask.priority,
                dueDate: editingTask.dueDate,
                status: editingTask.status,
            };

            await TaskService.updateTask(editingTask.ID, updatedTask);
            setEditingTask(null);  // Close the edit form
            fetchTasks();  // Reload tasks
        } catch (error) {
            console.error('Failed to update task:', error);
        }
    };

    useEffect(() => {
        fetchTasks();
    }, []);

    if (error) return <div>{error}</div>;

    return (
        <div>
            <h1>Task List</h1>
            <TaskForm onSuccess={fetchTasks} />
            {tasks.map((task) => (
                <div key={task.ID}>
                    <h3>{task.title}</h3>
                    <p>Task ID: {task.ID}</p>
                    <p>{task.description}</p>
                    <button onClick={() => handleEdit(task)}>Edit</button>
                    <button onClick={() => handleDelete(task.ID)}>Delete</button>
                </div>
            ))}

            {editingTask && (
                <form onSubmit={handleUpdate}>
                    <input
                        type="text"
                        value={editingTask.title}
                        onChange={(e) =>
                            setEditingTask({ ...editingTask, title: e.target.value })
                        }
                        required
                    />
                    <textarea
                        value={editingTask.description}
                        onChange={(e) =>
                            setEditingTask({ ...editingTask, description: e.target.value })
                        }
                    />
                    <button type="submit">Save</button>
                    <button onClick={() => setEditingTask(null)}>Cancel</button>
                </form>
            )}
        </div>
    );
};

export default TaskList;
