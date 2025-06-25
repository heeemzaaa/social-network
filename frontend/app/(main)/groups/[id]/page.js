export default async function Groups({params}) {
  let {id} = await params
  return (
    <div>Group id {id}</div>
  )
}
