export type Thread = {
    id: number;
    user_id: number;
    title: string;
    desc: string;
};

export type Post = {
    id: number;
    thread_id: number;
    user_id: number;
    title: string;
    body: string;
}

export type Comment = {
    id: number;
    post_id: number;
    user_id: number;
    body: string;
}