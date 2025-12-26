import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import type { Comment } from "../Type/models.ts"
import Ending from "../Components/CommentList.tsx";

function Comments() {
    const {id} = useParams<{ id: string }>();
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
                const lst = data.payload?.data ?? [];
                setComments(lst);
            } catch (err) {
                setError("Network or Server Error")
            }
        }
        load();
    }, [])
    
    if (error) return <>{error}</>;

    return (
        <>
            {comments.map((comment) => (
                <Ending
                    key={comment.id}
                    type="comment"
                    id={comment.id}
                    text={comment.body}
                />
            ))}
        </>
    );
}

export default Comments;