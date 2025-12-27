import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import type { Post, Thread } from "../Type/models.ts"
import Topic from "../Components/TopicList.tsx";
import { Card, CardContent, Typography } from "@mui/material";

function Posts() {
    const {id} = useParams<{ id: string }>();
    const [thread, setThread] = useState<Thread | null>(null);
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
                const { head, tail } = data.payload?.data ?? { head: null, tail: [] };
                setThread(head);
                setPosts(tail);
            } catch (err) {
                setError("Network or Server Error")
            }
        }
        load();
    }, [])
    
    if (error) return <>{error}</>;

    return (
        <Card>
            <CardContent>
                {thread && (
                    <>
                        <Typography variant="h4">
                            {thread.title}
                        </Typography>
                        <Typography variant="h6">
                            {thread.desc}
                        </Typography>
                    </>
                )}
                {posts.map((post) => (
                    <Topic
                        key={post.id}
                        type="post"
                        id={post.id}
                        title={post.title}
                        text={post.body}
                    />
                ))}
            </CardContent>
        </Card>
    );
}

export default Posts;