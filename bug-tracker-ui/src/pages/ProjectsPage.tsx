import { useCallback, useEffect, useState } from 'react';
import { FaPlus, FaTimes } from 'react-icons/fa';
import axiosInstance from '../api/axiosInstance';
import { toast } from 'react-toastify';
import type { Project } from '../types/types';
import { useNavigate } from 'react-router-dom';

const ProjectsPage = () => {
  const [projects, setProjects] = useState<Project[]>([]);
  const [search, setSearch] = useState('');
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const navigate = useNavigate();

  const fetchProjects = useCallback(async (searchTerm: string) => {
    try {
      const response = await axiosInstance.get('/project', {
        params: {
          search: searchTerm,
          page: 1,
          limit: 10,
        },
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      });

      setProjects(response.data.data);
    } catch (error) {
      console.error('Error fetching projects:', error);
    }
  }, []);

  useEffect(() => {
    fetchProjects(search);
  }, [fetchProjects, search]);

  const handleSearchBarChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(e.target.value);
  };

  const closeModal = () => {
    const modal = document.getElementById('addProjectModal') as HTMLDialogElement | null;
    modal?.close();
  };

  const openModal = () => {
    const modal = document.getElementById('addProjectModal') as HTMLDialogElement | null;
    modal?.showModal();
  };

  const handleSubmit = async () => {
    try {
      const response = await axiosInstance.post(
        '/project',
        { title: title, description: description },
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        }
      );

      if (response.data.success) {
        toast.success('Project added sucessfully!', {
          autoClose: 1000,
        });

        closeModal();
        fetchProjects('');
      }
    } catch (error: any) {
      console.error('Error adding project:', error);
      toast.error('Error adding project', {
        autoClose: 1000,
      });
    }
  };

  return (
    <div className="my-10 mx-20">
      {/* Header */}
      <div className="flex items-center justify-between mb-8">
        <div className="flex items-center gap-4">
          {' '}
          <h1 className="underline text-5xl font-bold">Projects</h1>
          <FaPlus className="text-2xl btn btn-sm btn-circle btn-ghost" onClick={() => openModal} />
        </div>
        <div className="">
          <input type="search" className="grow bg-light-gray px-4 py-1 rounded-lg" placeholder="Search" onChange={handleSearchBarChange} />
        </div>
      </div>

      {/* Projects List */}
      {/* TODO: refactor this to use same project list component that will be used on dashboard */}
      <div className="grid grid-cols-3 gap-10">
        {projects.map((project) => (
          <div key={project.id} className="bg-light-gray text-center p-20 hover:bg-gray-400" onClick={() => navigate(`/projects/${project.id}`)}>
            <h1 className="text-2xl font-bold">{project.title}</h1>
          </div>
        ))}
      </div>

      {/* Add Project Dialog */}
      <dialog id="addProjectModal" className="modal">
        <div className="modal-box bg-white">
          <FaTimes className=" btn btn-sm btn-circle btn-ghost absolute right-2 top-2" onClick={() => closeModal} />

          <h2 className="font-bold text-3xl text-center mt-4 mb-8">Add Project</h2>
          <div className="form-group">
            <label className="input-label">Title</label>
            <br />
            <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} required className="input-field" />
          </div>
          <div className="form-group">
            <label className="input-label">Description</label>
            <br />
            <textarea value={description} onChange={(e) => setDescription(e.target.value)} required className="input-field" rows={4} />
          </div>

          <div className="flex items-center gap-4 mx-20">
            <button onClick={() => handleSubmit()} className="button-base button-primary">
              Add
            </button>
            <button className="button-base button-secondary" onClick={() => closeModal}>
              Cancel
            </button>
          </div>
        </div>
      </dialog>
    </div>
  );
};

export default ProjectsPage;
