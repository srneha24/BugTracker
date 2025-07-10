import { useEffect, useState } from 'react';
import { FaPlus } from 'react-icons/fa';
import axiosInstance from '../api/axiosInstance';
import type { Project } from '../types/types';

const ProjectsPage = () => {
  const [projects, setProjects] = useState<Project[]>([]);

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const response = await axiosInstance.get('/project', {
          params: {
            search: '',
            page: 1,
            limit: 10,
          },
          headers: {
            Authorization: 'Bearer ' + localStorage.getItem('token'),
          },
        });

        setProjects(response.data.data);
      } catch (error) {
        console.error('Error fetching data: ', error);
      }
    };

    fetchProjects();
  }, []);

  return (
    <div className="my-10 mx-20">
      {/* Header */}
      <div className="flex items-center justify-between mb-8">
        <div className="flex items-center gap-4">
          {' '}
          <h1 className="underline text-5xl font-bold">Projects</h1>
          <FaPlus className="text-2xl" />
        </div>
        <div className="">
          <input type="search" className="grow bg-light-gray px-4 py-1 rounded-lg" placeholder="Search" />
        </div>
      </div>

      {/* Projects List */}
      {/* TODO: refactor this to use same project list component that will be used on dashboard */}
      <div className="grid grid-cols-3 gap-10">
        {projects.map((project) => (
          <div key={project.id} className="bg-light-gray text-center p-20 hover:bg-gray-400">
            <h1 className="text-2xl font-bold">{project.title}</h1>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ProjectsPage;
