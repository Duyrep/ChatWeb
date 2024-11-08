"use client";

import Message from "@/components/message";
import Loading from "@/components/loading";
import React, { useEffect, useState } from "react";

interface MessageInterface {
  _id: string;
  username: string;
  content: string;
}

export default function Chat() {
  const [messages, setMessages] = useState<MessageInterface[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  const sendMessage = (): void => {
    const messageInput = document.getElementById(
      "messageInput",
    ) as HTMLInputElement;
    const message = messageInput.value;
    fetch("http://127.0.0.1:8080/send_message", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ _id: null, username: "Client", content: message }),
    }).catch((error) => console.error(error));
    messageInput.value = "";
  };

  const getMessage = (): void => {
    fetch("http://127.0.0.1:8080/get_messages?amount=100")
      .then((response) => response.json())
      .then((data) => {
        if (data !== null) {
          setMessages(data);
        }
      })
      .catch((error) => console.error(error));
  };

  const handleScroll = (): void => {
    const maxScrollY =
      document.documentElement.scrollHeight - window.innerHeight;
    const scrollY = maxScrollY - window.scrollY;
    if (scrollY < 100 && !loading) {
      window.scrollTo(0, maxScrollY);
    }
  };

  useEffect(() => handleScroll(), [messages]);

  useEffect(() => {
    getMessage();
    setLoading(false);

    return () => {
      setLoading(true);
      setMessages([])
    };
  }, []);

  useEffect(() => {
    const ws = new WebSocket("ws://127.0.0.1:8080/ws");

    ws.addEventListener("message", (_) => {
      getMessage();
    });

    return () => {
      ws.close();
    };
  }, []);

  return (
    <div>
      {loading ? (
        <Loading />
      ) : (
        <div id="message-box" className="mt-14">
          <ul>
            {messages.map((message, index) => (
              <Message
                key={index}
                username={message.username}
                content={message.content}
              />
            ))}
          </ul>
        </div>
      )}

      <div className="fixed bottom-0 left-0 flex w-full border-t-2 border-yellow-500 bg-white p-2">
        <input
          className="w-full rounded-md border-2 border-black pl-2"
          id="messageInput"
          type="text"
          autoComplete="off"
          maxLength={200}
        />
        <button
          className="ml-2 rounded-md bg-black px-1 text-white"
          onClick={sendMessage}
        >
          send
        </button>
      </div>
    </div>
  );
}
