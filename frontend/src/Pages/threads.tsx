import { Link } from "react-router-dom";
import { useEffect, useState } from "react";

type Thread = {
    id: number;
    user_id: number;
    title: string;
    desc: string;
}

function Threads() {
    const [threads, setThreads] = useState<Thread[]>([]);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        async function load() {
            try {
                const res = await fetch("http://localhost:8000/threads", {
                    method: "GET",
                });

                if (!res.ok) {
                    setError(`Server Returned Status ${res.status}`);
                    return;
                }

                const data = await res.json();
                const lst = data.payload?.data ?? []
                setThreads(lst);
            } catch (err) {
                setError("Network or Server Error")
            }
        }
        load();
    }, [])
    
    if (error) return <>{error}</>;
    return (
        <div>
            {threads.map((thread) => (
                <Link
                    key={thread.id}
                    to={`/threads/${thread.id}`}
                    style={{ textDecoration: "none"}}>
                    <div key={thread.id} className="thread" >
                        <h2 className="thread-title">{thread.title}</h2>
                        <p className="thread-text">{thread.desc}</p>
                    </div>
                </Link>
            ))}
        </div>
    );
}

export default Threads;
// Send Cookie to Backend
// Login to Assign Cookie