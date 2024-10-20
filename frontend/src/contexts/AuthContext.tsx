import React, { createContext, useContext, useState, useEffect } from 'react';
import { User } from '../types/user';
import { AuthService } from '../services/authService';

interface AuthContextType {
    user: User | null;
    login: (email: string, password: string) => Promise<void>;
    logout: () => void;
    loading: boolean;
}

const AuthContext = createContext<AuthContextType>({} as AuthContextType);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            // Fetch user data from the backend if needed
            setUser({ id: 1, email: 'example@example.com', firstName: 'John', lastName: 'Doe', role: 'user' }); // Replace with actual API call
        }
        setLoading(false);
    }, []);

    const login = async (email: string, password: string) => {
        try{
            const token = await AuthService.login({ email, password });
            // Set the user data after login (API call for actual data)
            setUser({ id: 1, email, firstName: 'John', lastName: 'Doe', role: 'user' });
        }catch (error) {
            console.error('Login failed:', error);
            throw error;  // Ensure the error propagates to your component
        }
    };

    const logout = () => {
        AuthService.logout();
        setUser(null);
    };

    return (
        <AuthContext.Provider value={{ user, login, logout, loading }}>
            {!loading && children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => useContext(AuthContext);
