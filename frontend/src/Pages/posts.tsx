import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import type { Post } from "../Type/models.ts"
import Topic from "../Components/TopicList.tsx";

function Posts() {
    const {id} = useParams<{ id: string }>();
    const [posts, setPosts] = useState<Post[]>([]);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        async function load() {
            try {
                const res = await fetch(`http://localhost:8000/thread/${id}`, {
                    method: "GET",
                    credentials: "include",
                });

                if (!res.ok) {
                    setError(`Server Returned Status ${res.status}`);
                    return;
                }

                const data = await res.json();
                const lst = data.payload?.data ?? [];
                setPosts(lst);
            } catch (err) {
                setError("Network or Server Error")
            }
        }
        load();
    }, [])
    
    if (error) return <>{error}</>;

    return (
        <>
            {posts.map((post) => (
                <Topic
                    key={post.id}
                    type="post"
                    id={post.id}
                    title={post.title}
                    text={post.body}
                />
            ))}
        </>
    );
}

export default Posts;