
import React, {useState} from 'react';
import { Paper, Stack, Typography } from "@mui/material";
import {DndContext} from '@dnd-kit/core';
import {useDroppable} from '@dnd-kit/core';
import {useDraggable} from '@dnd-kit/core';


function Card(props) {
  const { attributes, listeners, setNodeRef } = useDraggable({
    id: props.id,
  })

  return (
    <Paper sx={{ margin: "6px" }} {...listeners} {...attributes} ref={setNodeRef}>
      {props.title}
    </Paper>
  )
}

function Swimlane({ title, items }) {
  const { setNodeRef } = useDroppable({id: title})
  return (
      <Stack direction="column" ref={setNodeRef} sx={{ width: "100%", height: "90vh", backgroundColor: '#C1F5E9', '&:hover': { backgroundColor: '#3AF4C9', opacity: [0.9, 0.8, 0.7] }}}>
        <Typography variant="h5" gutterBottom sx={{textAlign: 'center', textDecoration: "underline"}}>{title}</Typography>
        {items.map(({ title: cardTitle }, key) => {
          cardTitle = cardTitle ? cardTitle : "unset card title"
          return <Card title={cardTitle} key={key} index={key} parent={title} />
        })}
    </Stack>
  )
}


{/* Using https://blog.logrocket.com/build-kanban-board-dnd-kit-react/?ssp=1&darkschemeovr=1&setlang=en-AU&safesearch=moderate as the example how to do this */}

export default function Swimlanes() {
  const [todos, setTodos] = useState([<Card title="some-to-do" index="0" parent="" />, <Card title="some-to-do-2" index="1" parent="" />])
  const [inProgress, setInProgress] = useState([<Card title="some-in-progress" index="0" parent="" />])
  const [dones, setDones] = useState([<Card title="some-done" index="0" parent="" />])

  return (
    <Stack direction="row" spacing={2} justifyContent="space-between" alignItems="center">
      <DndContext>
        <Swimlane title="To Do" items={todos}/>
        <Swimlane title="In Progress" items={inProgress}/>
        <Swimlane title="Done" items={dones}/>
      </DndContext>
    </Stack>
  );

}