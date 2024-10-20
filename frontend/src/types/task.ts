export interface Task {
    ID: number;
    title: string;
    description: string;
    dueDate: string;
    priority: 'low' | 'medium' | 'high';
    status: string;
    userId: number;
}

export interface TaskCreate {
    title: string;
    description: string;
    dueDate: string;
    priority: 'low' | 'medium' | 'high';
}

export interface TaskUpdate {
    title?: string;
    description?: string;
    dueDate?: string;
    priority?: 'low' | 'medium' | 'high';
    status?: string;
}
