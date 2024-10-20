export interface User {
    id: number;
    email: string;
    firstName: string;
    lastName: string;
    role: string;
}

export interface UserLogin {
    email: string;
    password: string;
}

export interface UserRegister {
    email: string;
    password: string;
    firstName: string;
    lastName: string;
}
