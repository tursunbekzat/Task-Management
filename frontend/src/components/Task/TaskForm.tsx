import React, { useState } from 'react';
import { TaskCreate } from '../../types/task';
import { TaskService } from '../../services/taskService';

const TaskForm: React.FC<{ onSuccess?: () => void }> = ({ onSuccess }) => {
    const [formData, setFormData] = useState<TaskCreate>({
        title: '',
        description: '',
        dueDate: '',
        priority: 'medium',
    });

    const formatDueDate = (dueDate: string) => {
        // Ensure the datetime-local value is converted to ISO 8601 format
        const formattedDate = `${dueDate}:00Z`;
        return formattedDate;
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const formattedDueDate = formatDueDate(formData.dueDate);  // Format dueDate
            await TaskService.createTask({ ...formData, dueDate: formattedDueDate });
            setFormData({ title: '', description: '', dueDate: '', priority: 'medium' });
            onSuccess?.();  // Refresh task list or trigger success callback
        } catch (error) {
            console.error('Failed to create task:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                placeholder="Title"
                value={formData.title}
                onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                required
            />
            <textarea
                placeholder="Description"
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
            />
            <input
                type="datetime-local"
                value={formData.dueDate}
                onChange={(e) => setFormData({ ...formData, dueDate: e.target.value })}
                required
            />
            <select
                value={formData.priority}
                onChange={(e) => setFormData({ ...formData, priority: e.target.value as 'low' | 'medium' | 'high' })}
            >
                <option value="low">Low</option>
                <option value="medium">Medium</option>
                <option value="high">High</option>
            </select>
            <button type="submit">Create Task</button>
        </form>
    );
};

export default TaskForm;
