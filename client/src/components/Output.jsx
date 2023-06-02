import { useState } from 'react'
import './Output.css'

function Output({ outputText }) {
  const [copied, setCopied] = useState(false)

  const copyTextToClipboard = async (text) => {
      if ('clipboard' in navigator) {
        return await navigator.clipboard.writeText(text);
      } else {
        return document.execCommand('copy', true, text);
      }
    }

  const handleClick = (text) => {
    copyTextToClipboard(text).then(() => {
      setCopied(true)
      setTimeout(() => {
        setCopied(false)
      }, 2000)
    }).catch((err) => {
        console.err(err)
      })

  }
  return (
    <div className="output">
      <input type="text" value={outputText} readOnly />
      <button onClick={() => handleClick(outputText)}>
        <span>{copied ? 'Copied!' : 'Copy'}</span>
      </button>
    </div>
  );
}

export default Output
