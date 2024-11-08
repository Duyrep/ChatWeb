"use client"

import { useEffect, useState } from "react"

export default function Test() {
  const [num, setNum] = useState(10)

  useEffect(() => {
    console.log("Mount")
    setNum(1)

    return () => {
      console.log("Unmount")
    }
  }, [])
  return <div>Test</div>
}