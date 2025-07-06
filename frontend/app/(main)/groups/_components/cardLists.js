




export function CardList({ title, items }) {

    return (
        <div>
            {title && <h2>{title}</h2>}
            <ul>
                {items.map(item => (
                    <li key={item.id}> { }
                        <h3>{item.title}</h3>
                        <p>{item.description}</p>
                        <img src={`${item.image}`}></img>
           
                    </li>
                ))}
            </ul>
        </div>
    );

}


export const myItems = [
    { id: 1, title: 'Item A', description: 'Description for Item A', image:"path_to_image" },
    { id: 2, title: 'Item B', description: 'Description for Item B' },
    { id: 3, title: 'Item C', description: 'Description for Item C' },
];
