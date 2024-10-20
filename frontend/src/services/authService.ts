import axios from 'axios';
import { UserLogin, UserRegister } from '../types/user';

const API_URL = process.env.REACT_APP_API_URL;

export class AuthService {
    static async login(credentials: { email: string; password: string }): Promise<void> {
        try {
            const response = await axios.post(`${API_URL}/auth/login`, credentials);
            const token = response.data.token;
            localStorage.setItem('token', token);
            axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
        } catch (error) {
            console.error('Login failed:', error);
            throw error;
        }
    }

    static async register(userData: UserRegister): Promise<void> {
        await axios.post(`${API_URL}/auth/register`, userData);
    }

    static logout(): void {
        localStorage.removeItem('token');
        delete axios.defaults.headers.common['Authorization'];
    }

    static isAuthenticated(): boolean {
        return !!localStorage.getItem('token');
    }
}
