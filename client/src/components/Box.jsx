import { useState, useRef } from 'react'
import './Box.css'
import Output from './Output'
import Input from './Input'

function Box() {
  const [shortUrl, setShortUrl] = useState("")
  const inputRef = useRef(null)

  const handleGetShortUrl = async () => {
    const data = await fetch("http://localhost:3002/url", 
      { 
        method: "POST",
        body: JSON.stringify({ originalUrl: inputRef.current.value })
      }
    )
    const json = await data.json()
    setShortUrl(`http://localhost:3002/${json.shortUrl}`)
  }

  return(
    <div className='box'>
      <h1>Enter a URL to Shorten</h1> 
      <div className='controls'>
        <Input inputRef={inputRef} handleGetShortUrl={handleGetShortUrl} />
        <Output outputText={shortUrl} />
      </div>
    </div>
  )
}

export default Box
