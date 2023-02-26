
import React, {useState} from 'react';
import { Box, Paper, Stack, Typography } from "@mui/material";
import {DndContext} from '@dnd-kit/core';
import {useDroppable} from '@dnd-kit/core';
import {useDraggable} from '@dnd-kit/core';
import {CSS} from '@dnd-kit/utilities';

const Draggable = () => {
  const {attributes, listeners, setNodeRef} = useDraggable({
    id: 'draggable-1',
    data: {parent: 'ToDo', title: 'Complete blogpost.'}
  })

  return <div {...attributes} {...listeners} ref={setNodeRef}>Drag Me!</div>
}

const Droppable = () => {
  const {setNodeRef} = useDroppable({
    id: 'droppable-1'
  })

  return <div ref={setNodeRef}> Drop on me! </div>
}

const Card = (title, index, parent) => {
  const { attributes, listeners, setNodeRef } = useDraggable({
    id: title,
    data: {
      title,
      index,
      parent,
    },
  })

  return (
    <Box {...listeners} {...attributes} ref={setNodeRef}>
      {title}
    </Box>
  )
}

function Swimlane({ title, items }) {
  const { setNodeRef } = useDroppable({
    id: title,
  });

  return (
      <Box sx={{ width: "100%", height: "90vh", backgroundColor: '#C1F5E9', '&:hover': { backgroundColor: '#3AF4C9', opacity: [0.9, 0.8, 0.7] }}}>
        <Typography variant="h5" gutterBottom sx={{textAlign: 'center', textDecoration: "underline"}}>{title}</Typography>
        {items.map(({ title: cardTitle }, key) => (
          <Card title={cardTitle} key={key} index={key} parent={title} />
        ))}
    </Box>
  )
}


{/* Using https://blog.logrocket.com/build-kanban-board-dnd-kit-react/?ssp=1&darkschemeovr=1&setlang=en-AU&safesearch=moderate as the example how to do this */}

export default function Swimlanes() {
  const [todos, setTodos] = useState([<Card title="some-to-do" index="0" parent="" />])
  const [inProgress, setInProgress] = useState([<Card title="some-in-progress" index="0" parent="" />])
  const [dones, setDones] = useState([<Card title="some-done" index="0" parent="" />])

  return (
    <Stack direction="row" spacing={2} justifyContent="space-between" alignItems="center">
      <DndContext>
        <Swimlane title="To Do" items={todos}/>
        <Swimlane title="In Progress" items={inProgress}/>
        <Swimlane title="Done" itesm={dones}/>
      </DndContext>
    </Stack>
  );

}