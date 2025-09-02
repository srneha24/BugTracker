import React from "react";
import {
  useQuery,
  useQueryClient,
  keepPreviousData,
} from "@tanstack/react-query";
import { motion } from "framer-motion";
import {
  AlertCircle,
  ArrowUpDown,
  Loader2,
  RefreshCw,
  ChevronRight,
  Filter,
  Search,
  X,
} from "lucide-react";

// ------------------ Types ------------------
export type Bug = {
  id: string;
  title: string;
  project: string;
  status: string;
  priority: string;
  created_at?: string;
};
export type Project = { id: string; name: string };

// ------------------ API utils ------------------
const API_URL =
  (typeof import.meta !== "undefined" &&
    (import.meta as any).env?.VITE_API_URL) ||
  "http://localhost:8080";

const cn = (...s: Array<string | false | undefined>) =>
  s.filter(Boolean).join(" ");

function authHeaders(): HeadersInit {
  const token =
    typeof window !== "undefined" ? localStorage.getItem("authToken") : null;
  const h: Record<string, string> = { "Content-Type": "application/json" };
  if (token) h.Authorization = `Bearer ${token}`;
  return h;
}
async function apiGet<T>(path: string, signal?: AbortSignal): Promise<T> {
  const res = await fetch(`${API_URL}${path}`, {
    method: "GET",
    headers: authHeaders(),
    credentials: "include",
    signal,
  });
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  return (await res.json()) as T;
}

// ------------------ Adapters ------------------
type BugDto = {
  id: string | number;
  title: string;
  status: string;
  priority: string | number;
  project?: string | { name?: string };
  project_name?: string;
  created_at?: string;
};
function adaptBug(d: BugDto): Bug {
  const project = typeof d.project === "string" ? d.project : d.project?.name;
  return {
    id: String(d.id),
    title: d.title,
    status: d.status,
    priority: String(d.priority).toUpperCase(),
    project: d.project_name || project || "—",
    created_at: d.created_at,
  };
}
type ProjectDto = { id: string | number; name?: string; title?: string };
function adaptProject(d: ProjectDto): Project {
  return { id: String(d.id), name: d.name || d.title || "Untitled" };
}

// ------------------ Fetchers ------------------
type BugsParams = {
  q?: string;
  status?: string;
  priority?: string;
  sort?: string;
  limit?: number;
  offset?: number;
};
async function fetchBugs(
  p: BugsParams,
  signal?: AbortSignal
): Promise<{ items: Bug[]; total: number }> {
  const usp = new URLSearchParams();
  if (p.q) usp.set("q", p.q);
  if (p.status) usp.set("status", p.status);
  if (p.priority) usp.set("priority", p.priority);
  if (p.sort) usp.set("sort", p.sort);
  usp.set("limit", String(p.limit ?? 20));
  usp.set("offset", String(p.offset ?? 0));
  const data = await apiGet<{ items: BugDto[]; total?: number }>(
    `/api/bugs?${usp.toString()}`,
    signal
  );
  return {
    items: (data.items || []).map(adaptBug),
    total: data.total ?? (data.items?.length || 0),
  };
}
async function fetchProjects(
  limit = 6,
  offset = 0,
  signal?: AbortSignal
): Promise<{ items: Project[]; total: number }> {
  const data = await apiGet<{ items: ProjectDto[]; total?: number }>(
    `/api/projects?limit=${limit}&offset=${offset}`,
    signal
  );
  return {
    items: (data.items || []).map(adaptProject),
    total: data.total ?? (data.items?.length || 0),
  };
}

// ------------------ Visual helpers ------------------
const statusTone: Record<string, string> = {
  New: "bg-blue-100 text-blue-800 border-blue-200",
  "In Progress": "bg-amber-100 text-amber-800 border-amber-200",
  Resolved: "bg-emerald-100 text-emerald-800 border-emerald-200",
  Closed: "bg-zinc-100 text-zinc-800 border-zinc-200",
};
const priorityTone: Record<string, string> = {
  P1: "bg-red-100 text-red-700",
  P2: "bg-orange-100 text-orange-700",
  P3: "bg-yellow-100 text-yellow-800",
  P4: "bg-zinc-100 text-zinc-700",
};

// ------------------ Small UI ------------------
function HeaderBar() {
  return (
    <header className="sticky top-0 z-20 border-b bg-white/80 backdrop-blur">
      <div className="mx-auto flex max-w-7xl items-center justify-between px-6 py-3">
        <div className="flex items-center gap-3">
          <div className="grid h-7 w-20 place-items-center rounded-md bg-black text-xs font-bold tracking-wide text-white">
            LOGO
          </div>
          <span className="hidden text-sm text-gray-500 sm:inline-flex">
            Bug Tracker
          </span>
        </div>
        <div className="flex items-center gap-3">
          <button className="rounded-md border px-3 py-1.5 text-sm hover:bg-gray-50">
            Projects
          </button>
          <div className="grid h-8 w-8 place-items-center rounded-full bg-gray-200 text-xs font-semibold">
            BR
          </div>
        </div>
      </div>
    </header>
  );
}

function Toolbar(props: {
  q: string;
  setQ: (v: string) => void;
  status: string;
  setStatus: (v: string) => void;
  priority: string;
  setPriority: (v: string) => void;
  sort: string;
  setSort: (v: string) => void;
  onClear: () => void;
  onRefresh: () => void;
  refreshing: boolean;
}) {
  const {
    q,
    setQ,
    status,
    setStatus,
    priority,
    setPriority,
    sort,
    setSort,
    onClear,
    onRefresh,
    refreshing,
  } = props;
  return (
    <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
      <div className="flex w-full flex-1 items-center gap-2">
        <div className="relative w-full max-w-sm">
          <Search className="pointer-events-none absolute left-2 top-1/2 h-4 w-4 -translate-y-1/2 text-gray-400" />
          <input
            value={q}
            onChange={(e) => setQ(e.target.value)}
            placeholder="Search bugs…"
            className="w-full rounded-md border px-8 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-black/10"
          />
          {q && (
            <button
              aria-label="Clear search"
              className="absolute right-2 top-1/2 -translate-y-1/2 text-gray-500"
              onClick={() => setQ("")}
            >
              <X className="h-4 w-4" />
            </button>
          )}
        </div>
        <button
          className="grid h-9 w-9 place-items-center rounded-md border"
          onClick={onRefresh}
          disabled={refreshing}
          title="Refresh"
        >
          {refreshing ? (
            <Loader2 className="h-4 w-4 animate-spin" />
          ) : (
            <RefreshCw className="h-4 w-4" />
          )}
        </button>
      </div>

      <div className="flex items-center gap-2">
        <select
          value={status}
          onChange={(e) => setStatus(e.target.value)}
          className="rounded-md border px-2 py-2 text-sm"
        >
          <option value="">All statuses</option>
          <option value="active">Active</option>
          <option value="New">New</option>
          <option value="In Progress">In Progress</option>
          <option value="Resolved">Resolved</option>
          <option value="Closed">Closed</option>
        </select>
        <select
          value={priority}
          onChange={(e) => setPriority(e.target.value)}
          className="rounded-md border px-2 py-2 text-sm"
        >
          <option value="">All priorities</option>
          <option value="P1">P1 – Critical</option>
          <option value="P2">P2 – High</option>
          <option value="P3">P3 – Medium</option>
          <option value="P4">P4 – Low</option>
        </select>
        <select
          value={sort}
          onChange={(e) => setSort(e.target.value)}
          className="rounded-md border px-2 py-2 text-sm"
        >
          <option value="created_at:desc">Newest</option>
          <option value="created_at:asc">Oldest</option>
          <option value="priority:asc">Priority ↑</option>
          <option value="priority:desc">Priority ↓</option>
          <option value="title:asc">Title A→Z</option>
          <option value="title:desc">Title Z→A</option>
        </select>
        <button
          className="inline-flex items-center gap-2 rounded-md px-3 py-2 text-sm text-gray-600 hover:bg-gray-100"
          onClick={onClear}
        >
          <Filter className="h-4 w-4" />
          Clear
        </button>
      </div>
    </div>
  );
}

function BugRow({ bug, index }: { bug: Bug; index: number }) {
  return (
    <tr
      className={cn(
        "border-t transition hover:bg-gray-50",
        index % 2 ? "bg-white" : "bg-gray-50/50"
      )}
    >
      <td className="px-5 py-4 font-medium">{bug.title}</td>
      <td className="px-5 py-4 text-gray-500">{bug.project}</td>
      <td className="px-5 py-4">
        <span
          className={cn(
            "inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-medium",
            statusTone[bug.status] || "bg-gray-100 text-gray-700 border-gray-200"
          )}
        >
          {bug.status}
        </span>
      </td>
      <td className="px-5 py-4">
        <span
          className={cn(
            "inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-semibold",
            priorityTone[bug.priority] || "bg-gray-100 text-gray-700"
          )}
        >
          {bug.priority}
        </span>
      </td>
      <td className="whitespace-nowrap px-5 py-4 text-sm text-gray-500">
        {bug.created_at ? new Date(bug.created_at).toLocaleDateString() : "—"}
      </td>
      <td className="px-5 py-4 text-right">
        <button className="inline-flex items-center gap-1 rounded-md px-2 py-1 text-sm text-gray-600 hover:bg-gray-100">
          Open <ChevronRight className="h-4 w-4" />
        </button>
      </td>
    </tr>
  );
}
function SkeletonRow() {
  return (
    <tr className="animate-pulse border-t">
      {Array.from({ length: 6 }).map((_, i) => (
        <td key={i} className="px-5 py-4">
          <div className="h-4 w-28 rounded bg-gray-200" />
        </td>
      ))}
    </tr>
  );
}
function BugsTable(props: {
  bugs: Bug[];
  loading: boolean;
  error: boolean;
  onRetry: () => void;
  page: number;
  pageSize: number;
  total: number;
  onPageChange: (p: number) => void;
}) {
  const { bugs, loading, error, onRetry, page, pageSize, total, onPageChange } =
    props;
  const totalPages = Math.max(1, Math.ceil(total / pageSize));
  return (
    <div className="overflow-hidden rounded-2xl border">
      <table className="w-full text-left">
        <thead>
          <tr className="bg-gray-100 text-gray-600">
            <th className="px-5 py-3 text-sm font-medium">Title</th>
            <th className="px-5 py-3 text-sm font-medium">Project</th>
            <th className="px-5 py-3 text-sm font-medium">Status</th>
            <th className="px-5 py-3 text-sm font-medium">Priority</th>
            <th className="px-5 py-3 text-sm font-medium">Created</th>
            <th className="px-5 py-3 text-sm font-medium text-right"></th>
          </tr>
        </thead>
        <tbody>
          {loading && (
            <>
              <SkeletonRow />
              <SkeletonRow />
              <SkeletonRow />
              <SkeletonRow />
            </>
          )}
          {!loading && error && (
            <tr>
              <td colSpan={6} className="px-5 py-6 text-sm text-red-600">
                <div className="flex items-center gap-2">
                  <AlertCircle className="h-4 w-4" />
                  Failed to load bugs.{" "}
                  <button className="underline" onClick={onRetry}>
                    Retry
                  </button>
                </div>
              </td>
            </tr>
          )}
          {!loading && !error &&
            bugs.map((b, i) => <BugRow key={b.id} bug={b} index={i} />)}
          {!loading && !error && bugs.length === 0 && (
            <tr>
              <td
                colSpan={6}
                className="px-5 py-10 text-center text-sm text-gray-500"
              >
                No bugs match your filters.
              </td>
            </tr>
          )}
        </tbody>
      </table>

      <div className="flex items-center justify-between px-5 py-3">
        <div className="text-xs text-gray-500">
          Showing {(page - 1) * pageSize + (bugs.length ? 1 : 0)}–
          {(page - 1) * pageSize + bugs.length} of {total}
        </div>
        <div className="flex items-center gap-2">
          <button
            className="rounded-md border px-2 py-1 text-sm disabled:opacity-50"
            onClick={() => onPageChange(Math.max(1, page - 1))}
            disabled={page <= 1}
          >
            Prev
          </button>
          <span className="rounded-md bg-gray-100 px-2 py-1 text-xs">
            Page {page} / {Math.max(1, totalPages)}
          </span>
          <button
            className="rounded-md border px-2 py-1 text-sm disabled:opacity-50"
            onClick={() => onPageChange(Math.min(totalPages, page + 1))}
            disabled={page >= totalPages}
          >
            Next
          </button>
        </div>
      </div>
    </div>
  );
}

function ProjectsGrid(props: {
  data: Project[];
  loading: boolean;
  error: boolean;
  onRetry: () => void;
}) {
  const { data, loading, error, onRetry } = props;
  if (loading) {
    return (
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
        {Array.from({ length: 3 }).map((_, i) => (
          <div key={i} className="rounded-2xl border">
            <div className="grid h-40 place-items-center">
              <div className="h-6 w-28 animate-pulse rounded bg-gray-200" />
            </div>
          </div>
        ))}
      </div>
    );
  }
  if (error) {
    return (
      <div className="text-sm text-red-600">
        Failed to load projects.{" "}
        <button className="underline" onClick={onRetry}>
          Retry
        </button>
      </div>
    );
  }
  return (
    <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
      {data.slice(0, 3).map((p, idx) => (
        <motion.div
          key={p.id}
          initial={{ opacity: 0, y: 8 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.2, delay: idx * 0.05 }}
        >
          <div className="rounded-2xl border shadow-sm">
            <div className="border-b p-4 text-lg font-semibold">{p.name}</div>
            <div className="flex items-center justify-between p-4 text-sm text-gray-500">
              <div>Open issues • —</div>
              <button className="inline-flex items-center gap-1 rounded-md px-2 py-1 text-sm text-gray-600 hover:bg-gray-100">
                Open <ChevronRight className="h-4 w-4" />
              </button>
            </div>
          </div>
        </motion.div>
      ))}
    </div>
  );
}

// ------------------ Page ------------------
export default function HomePage() {
  const qc = useQueryClient();

  const [q, setQ] = React.useState("");
  const [status, setStatus] = React.useState("active");
  const [priority, setPriority] = React.useState("");
  const [sort, setSort] = React.useState("created_at:desc");
  const [page, setPage] = React.useState(1);
  const pageSize = 20;

  const {
    data: bugsData,
    isLoading: bugsLoading,
    error: bugsError,
    refetch: refetchBugs,
    isFetching: bugsFetching,
  } = useQuery<{ items: Bug[]; total: number }, Error>({
    queryKey: ["bugs", { q, status, priority, sort, page, pageSize }],
    queryFn: ({ signal }) =>
      fetchBugs(
        { q, status, priority, sort, limit: pageSize, offset: (page - 1) * pageSize },
        signal
      ),
    placeholderData: keepPreviousData,
  });

  const {
    data: projectsData,
    isLoading: projectsLoading,
    error: projectsError,
    refetch: refetchProjects,
  } = useQuery<{ items: Project[]; total: number }, Error>({
    queryKey: ["projects", 6, 0],
    queryFn: ({ signal }) => fetchProjects(6, 0, signal),
  });

  const onClear = () => {
    setQ("");
    setStatus("active");
    setPriority("");
    setSort("created_at:desc");
    setPage(1);
  };
  const onRefresh = () => {
    qc.invalidateQueries({ queryKey: ["bugs"] });
    qc.invalidateQueries({ queryKey: ["projects"] });
  };

  return (
    <div className="min-h-screen bg-white text-gray-900">
      <HeaderBar />

      <main className="mx-auto max-w-7xl px-6 pb-16 pt-8">
        <div className="mb-6 flex items-center justify-between">
          <h1 className="text-4xl font-extrabold tracking-tight">Welcome, USER</h1>
        </div>

        {/* Bugs */}
        <section className="mb-10">
          <div className="mb-4 flex items-center justify-between">
            <div className="flex items-center gap-2">
              <h2 className="text-2xl font-bold tracking-tight">Active Bugs</h2>
              <span className="inline-flex items-center gap-1 rounded-md bg-gray-100 px-2 py-1 text-xs">
                <ArrowUpDown className="h-3.5 w-3.5" />
                Sorted: {sort.split(":")[0]}
              </span>
            </div>
          </div>

          <Toolbar
            q={q}
            setQ={setQ}
            status={status}
            setStatus={(v) => {
              setStatus(v);
              setPage(1);
            }}
            priority={priority}
            setPriority={(v) => {
              setPriority(v);
              setPage(1);
            }}
            sort={sort}
            setSort={(v) => {
              setSort(v);
              setPage(1);
            }}
            onClear={onClear}
            onRefresh={onRefresh}
            refreshing={bugsFetching}
          />

          <div className="mt-4" />

          <BugsTable
            bugs={bugsData?.items || []}
            loading={bugsLoading}
            error={!!bugsError}
            onRetry={() => refetchBugs()}
            page={page}
            pageSize={pageSize}
            total={bugsData?.total || 0}
            onPageChange={setPage}
          />
        </section>

        <hr className="my-10 border-gray-200" />

        {/* Projects */}
        <section>
          <div className="mb-4 flex items-center justify-between">
            <h2 className="text-2xl font-bold tracking-tight">Projects</h2>
          </div>
          <ProjectsGrid
            data={projectsData?.items || []}
            loading={projectsLoading}
            error={!!projectsError}
            onRetry={() => refetchProjects()}
          />
        </section>
      </main>
    </div>
  );
}
