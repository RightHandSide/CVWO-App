import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import type { Post, Comment } from "../Type/models.ts"
import Ending from "../Components/CommentList.tsx";
import { Card, CardContent, Typography } from "@mui/material";

function Comments() {
    const {id} = useParams<{ id: string }>();
    const [post, setPost] = useState<Post | null>(null);
    const [comments, setComments] = useState<Comment[]>([]);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        async function load() {
            try {
                const res = await fetch(`http://localhost:8000/post/${id}`, {
                    method: "GET",
                    credentials: "include",
                });

                if (!res.ok) {
                    setError(`Server Returned Status ${res.status}`);
                    return;
                }

                const data = await res.json();
                const { head, tail } = data.payload?.data ?? { head: null, tail: [] };
                setPost(head);
                setComments(tail);
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
                {post && (
                    <>
                        <Typography variant="h4">
                            {post.title}
                        </Typography>
                        <Typography variant="h6">
                            {post.body}
                        </Typography>
                    </>
                )}
                {comments.map((comment) => (
                    <Ending
                        key={comment.id}
                        type="comment"
                        id={comment.id}
                        text={comment.body}
                    />
                ))}
            </CardContent>
        </Card>
    );
}

export default Comments;