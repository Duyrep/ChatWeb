export type MessageProps = {
  username: string;
  content: string;
};

export default function Message({ username, content }: MessageProps) {
  return (
    <div className="mt-2 flex">
      <img
        className="h-10 select-none rounded-3xl"
        src="https://cdn.discordapp.com/embed/avatars/0.png"
        alt="avatar"
      />
      <div className="ml-2 w-[100%-2.5rem] rounded-md border-2 border-yellow-500 p-2">
        <b>{username}</b>
        <p className="break-words">{content}</p>
      </div>
    </div>
  );
}
