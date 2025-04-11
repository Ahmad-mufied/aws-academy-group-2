export interface User {
  id: number; 
  name: string;
  email: string;
  dob: Date; 
  role: 'Admin' | 'Manager' | 'Supervisor' | 'Staff';
  registeredDate: Date; 
  status: 'Active' | 'Inactive';
  products: string[]; 
}
