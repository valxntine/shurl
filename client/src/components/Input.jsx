import './Input.css'

function Input({ inputRef, handleGetShortUrl }) {
  return(
    <div className="input">
      <input ref={inputRef} type='text' placeholder='e.g. https://www.cronofy.com/about'></input>
      <button onClick={handleGetShortUrl}>Shorten!</button>
    </div>
  )
}

export default Input
