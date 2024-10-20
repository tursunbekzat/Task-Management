import axios from 'axios';
import { Task, TaskCreate, TaskUpdate } from '../types/task';

const API_URL = process.env.REACT_APP_API_URL;

// Helper to get the token from localStorage
const getAuthToken = () => localStorage.getItem('token');

export const TaskService = {
    async listTasks(): Promise<Task[]> {
        const token = getAuthToken();  // Retrieve token from storage
        const response = await axios.get(`${API_URL}/tasks`, {
            headers: {
                Authorization: `Bearer ${token}`
            },
            withCredentials: true,
        });
        return response.data;
    },

    async createTask(task: TaskCreate): Promise<Task> {
        const token = getAuthToken();
        const response = await axios.post(`${API_URL}/tasks`, task, {
            headers: {
                Authorization: `Bearer ${token}`
            },
            withCredentials: true,
        });
        return response.data;
    },

    async updateTask(id: number, task: TaskUpdate): Promise<Task> {
        const token = getAuthToken();
        const response = await axios.put(`${API_URL}/tasks/${id}`, task, {
            headers: {
                Authorization: `Bearer ${token}`
            },
            withCredentials: true,
        });
        return response.data;
    },

    async deleteTask(id: number): Promise<void> {
        const token = getAuthToken();
        await axios.delete(`${API_URL}/tasks/${id}`, {
            headers: {
                Authorization: `Bearer ${token}`
            },
            withCredentials: true,
        });
    }
};
