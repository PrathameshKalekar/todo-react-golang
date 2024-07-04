import axios from 'axios'
import { ADD_TODO } from './component/utils/ApiRoutes.js'
import TODOS from './component/pages/TODOS.jsx';
import { useState } from 'react';

function App() {
  const [todoText, setTodoText] = useState("")
  const [refresh, setRefresh] = useState(false)

  const addTODO = async () => {
    if (todoText.length === 0) {
      alert("No text present to add in TODO")
    } else {

      try {
        const newTODO = {
          todo_id: Date.now().toString(),
          todo: todoText,
          is_done: false
        }
        const { data } = await axios.post(ADD_TODO, newTODO)
        setTodoText("")
        setRefresh(!refresh)
      } catch (error) {
        console.error(error)
      }
    }
  }

  return (
    <div className="flex justify-center items-center flex-col">
      <div className="justify-center bg-slate-400 flex w-full ">
        TODO
      </div>
      <div className="mt-4 flex flex-row">
        <input
          type="text"
          name="Enter TODO"
          id="TODO"
          value={todoText}
          onChange={(e) => { setTodoText(e.target.value) }}
          className="p-2 border border-gray-300 rounded w-64"
          placeholder="Enter your TODO"
        />
        <div className="flex mx-5 border cursor-pointer bg-slate-400 w-20 items-center justify-center rounded-lg hover:bg-slate-500" onClick={addTODO}>Add</div>
      </div>

      <div className="my-4 flex justify-center bg-slate-400 w-1/2 rounded-lg p-2">TODO List</div>

      <TODOS refresh={refresh} />

    </div>


  );
}

export default App;
