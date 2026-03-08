export type BugStatus = 'todo' | 'in_progress' | 'done';
export type Priority = 1 | 2 | 3;

export interface UserResponse {
  id: number;
  name: string;
  username: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface UserBugsProject {
  id: number;
  title: string;
}

export interface UserBugsResponse {
  id: number;
  title: string;
  status: BugStatus;
  priority: Priority;
  created_at: string;
  updated_at: string;
  project: UserBugsProject;
}
