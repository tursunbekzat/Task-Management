import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './components/Auth/Login';
import TaskList from './components/Task/TaskList';
import { AuthProvider } from './contexts/AuthContext';

const App: React.FC = () => (
    <AuthProvider>
        <Router>
            <Routes>
                <Route path="/login" element={<Login />} />
                <Route path="/tasks" element={<TaskList />} />
            </Routes>
        </Router>
    </AuthProvider>
);

export default App;
