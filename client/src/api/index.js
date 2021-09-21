const serverAddress =
  process.env.NODE_ENV === "production" ? "" : "http://localhost"

export async function newPost(text) {
  let resp = await fetch(`${serverAddress}/post`, {
    method: "POST",
    body: text
  })
  if (!resp.ok) {
    alert("Failed to add new post to the wall...")
  }
}

export async function fetchPosts(slice) {
  let resp = await fetch(`${serverAddress}/fetch/${slice}`)
  if (!resp.ok) {
    alert("Failed to fetch posts.")
    return []
  }
  let posts = await resp.json()
  return posts
}
