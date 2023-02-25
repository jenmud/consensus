
import React, {useState} from 'react';
import { Box, Paper, Stack } from "@mui/material";
import {DndContext} from '@dnd-kit/core';
import {useDroppable} from '@dnd-kit/core';
import {useDraggable} from '@dnd-kit/core';
import {CSS} from '@dnd-kit/utilities';

function Droppable(props) {
  const {isOver, setNodeRef} = useDroppable({
    id: props.id,
  });
  const style = {
    opacity: isOver ? 1 : 0.5,
  };

  return (
    <div ref={setNodeRef} style={style}>
      {props.children}
    </div>
  );
}

function Draggable(props) {
  const {attributes, listeners, setNodeRef, transform} = useDraggable({
    id: props.id,
  });
  const style = {
    // Outputs `translate3d(x, y, 0)`
    transform: CSS.Translate.toString(transform),
  };

  return (
    <button ref={setNodeRef} style={style} {...listeners} {...attributes}>
      {props.children}
    </button>
  );
}

export default function Swimlane() {
  return (
    <Stack direction="row" spacing={2} justifyContent="space-between" alignItems="center">
      <Box sx={{ width: "100%", height: 600, backgroundColor: 'primary.dark', '&:hover': { backgroundColor: 'primary.main', opacity: [0.9, 0.8, 0.7] }}}>
        swimlane 1
      </Box>

      <Box sx={{ width: "100%", height: 600, backgroundColor: 'primary.dark', '&:hover': { backgroundColor: 'primary.main', opacity: [0.9, 0.8, 0.7] }}}>
        swimlane 2
      </Box>

      <Box sx={{ width: "100%", height: 600, backgroundColor: 'primary.dark', '&:hover': { backgroundColor: 'primary.main', opacity: [0.9, 0.8, 0.7] }}}>
        swimlane 3
      </Box>
    </Stack>
  );

}