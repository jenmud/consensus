
import React, { useState } from 'react';
import { Box, Paper, Stack, Typography } from "@mui/material";
import { DndContext } from '@dnd-kit/core';
import { useDroppable } from '@dnd-kit/core';
import { useDraggable } from '@dnd-kit/core';
import { CSS } from '@dnd-kit/utilities';


function Card(props) {
  const { title, index, parent } = props
  const { attributes, listeners, setNodeRef, transform } = useDraggable({
    id: title,
    data: { title, index, parent },
  })

  const style = {
    // Outputs `translate3d(x, y, 0)`
    transform: CSS.Translate.toString(transform),
  };

  return (
    <Paper sx={{ margin: "6px" }} {...listeners} {...attributes} style={style} ref={setNodeRef}>
      {title}
    </Paper>
  )
}

function Swimlane({ title, items }) {
  const { setNodeRef } = useDroppable({ id: title })
  return (
    <Box ref={setNodeRef} sx={{ width: "100%", height: "90vh", backgroundColor: '#C1F5E9', '&:hover': { backgroundColor: '#3AF4C9', opacity: [0.9, 0.8, 0.7] } }}>
      <Typography variant="h5" gutterBottom sx={{ textAlign: 'center', textDecoration: "underline" }}>{title}</Typography>
      <Stack>
        {items.map(({ title: cardTitle }, key) => <Card title={cardTitle ? cardTitle : "unset"} key={key} index={key} parent={title} />)}
      </Stack>
    </Box>
  )
}


{/* Using https://blog.logrocket.com/build-kanban-board-dnd-kit-react/?ssp=1&darkschemeovr=1&setlang=en-AU&safesearch=moderate as the example how to do this */ }

export default function Swimlanes() {
  const [todos, setTodos] = useState([<Card title="some-to-do" index="0" parent="" />, <Card title="some-to-do-2" index="1" parent="" />])
  const [inProgress, setInProgress] = useState([<Card title="some-in-progress" index="0" parent="" />])
  const [dones, setDones] = useState([<Card title="some-done" index="0" parent="" />])

  return (
    <Stack direction="row" spacing={2} justifyContent="space-between" alignItems="center">
      <DndContext
        onDragEnd={(e) => {
          const container = e.over?.id;
          const title = e.active.data.current?.title ?? "";
          const index = e.active.data.current?.index ?? 0;
          const parent = e.active.data.current?.parent ?? "To Do";

          if (container === "To Do") {
            setTodos([...todos, { title }]);
          } else if (container === "In Progress") {
            setInProgress([...inProgress, { title }]);
          } else if (container === "Done") {
            setDones([...dones, { title }]);
          }

          if (parent === "To Do") {
            setTodos([
              ...todos.slice(0, index),
              ...todos.slice(index + 1),
            ]);
          } else if (parent === "In Progress") {
            setInProgress([
              ...inProgress.slice(0, index),
              ...inProgress.slice(index + 1),
            ]);
          } else if (parent === "Done") {
            setDones([
              ...dones.slice(0, index),
              ...dones.slice(index + 1),
            ]);
          }

        }}
      >
        <Swimlane title="To Do" items={todos} />
        <Swimlane title="In Progress" items={inProgress} />
        <Swimlane title="Done" items={dones} />
      </DndContext>
    </Stack>
  );

}