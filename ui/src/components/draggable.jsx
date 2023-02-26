
import React, {useState} from 'react';
import { Box, Paper, Stack, Typography } from "@mui/material";
import {DndContext} from '@dnd-kit/core';
import {useDroppable} from '@dnd-kit/core';
import {useDraggable} from '@dnd-kit/core';
import {CSS} from '@dnd-kit/utilities';


function Droppable(props) {
  const {isOver, setNodeRef} = useDroppable({id: 'droppable'});
  const style = {color: isOver ? 'green' : undefined,};
  return (
    <div ref={setNodeRef} style={style}>
      {props.children}
    </div>
  );
}

function Draggable(props) {
  const {attributes, listeners, setNodeRef, transform} = useDraggable({id: 'draggable'});
  const style = transform ? {transform: `translate3d(${transform.x}px, ${transform.y}px, 0)`} : undefined;
  return (
    <button ref={setNodeRef} style={style} {...listeners} {...attributes}>
      {props.children}
      some text {props.name}
    </button>
  );
}

function Card(props) {
  return (
    <Box>
          <Draggable name={props.name}/>
          <Droppable name={props.name} />
    </Box>
  )
}

function Swimlane(props) {

  return(
      <Box sx={{ width: "100%", height: "90vh", backgroundColor: '#C1F5E9', '&:hover': { backgroundColor: '#3AF4C9', opacity: [0.9, 0.8, 0.7] }}}>
        <Typography variant="h5" gutterBottom sx={{textAlign: 'center', textDecoration: "underline"}}>{props.kind}</Typography>
        <DndContext>
          <DndContext>
            <Card name="A"/>
          </DndContext>
          <DndContext>
            <Card name="B"/>
          </DndContext>
        </DndContext>
      </Box>
  )
}

export default function Swimlanes() {
  return (
    <Stack direction="row" spacing={2} justifyContent="space-between" alignItems="center">
      <Swimlane kind="To Do"/>
      <Swimlane kind="In Progress"/>
      <Swimlane kind="Done"/>
    </Stack>
  );

}