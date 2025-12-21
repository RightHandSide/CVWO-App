import { useEffect, useState } from "react";
import Topic from "../Component/Topic.tsx";

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
                    credentials: "include",
                });

                if (!res.ok) {
                    setError(`Server Returned Status ${res.status}`);
                    return;
                }

                const data = await res.json();
                const lst = data.payload?.data ?? [];
                setThreads(lst);
            } catch (err) {
                setError("Network or Server Error")
            }
        }
        load();
    }, [])
    
    if (error) return <>{error}</>;
    return (
        <>
            {threads.map((thread) => (
                <Topic type="threads" id={thread.id} title={thread.title} text={thread.desc} />
            ))}
        </>
    );
}

export default Threads;