import axios from "axios";
import { useEffect, useState } from "react";
import { DELETE_TODO, RETRIVE_ALL_TODOS, UPDATE_TODO_DONE } from "../utils/ApiRoutes.js";
import { FaTrashCan } from "react-icons/fa6";
import { MdCheckBoxOutlineBlank, MdOutlineCheckBox } from "react-icons/md";
function TODOS({ refresh }) {
    const [TODOS, setTODOS] = useState([])
    const [reload, setReload] = useState(false)

    const deleteTODO = async (id) => {
        try {
            const { data } = await axios.delete(`${DELETE_TODO}/${id}`)
            console.log(data);


        } catch (error) {
            console.error(error);
        }
    }

    const updateTODODone = async (id, done) => {
        try {
            if (done) {
            } else {

                const { data } = await axios.put(`${UPDATE_TODO_DONE}/${id}`)
                console.log(data);
                setReload(!reload)
            }
        } catch (error) {
            console.error(error);
        }
    }

    useEffect(() => {
        async function fetchTODOS() {
            console.log("retriving");
            try {
                const { data } = await axios.get(RETRIVE_ALL_TODOS)
                console.log(data);
                if (data.length > 0) {
                    setTODOS(data)
                } else {
                    setTODOS([])
                }
            } catch (error) {
                console.log("catch");
                console.error('Error fetching todos:', error);
            }
        }
        fetchTODOS()
    }, [refresh, reload])



    return (
        <div>
            {
                TODOS.length === 0 ? (
                    <div>No TODOS Available </div>
                ) : (
                    TODOS.map(todo => (
                        <div className="flex justify-start">
                            <div key={todo.todo_id} className={`flex my-2 py-2 border justify-center  sm:w-72 md:w-80 lg:w-96 rounded-lg mb-2 font-bold text-slate-200 ${todo.is_done ? "bg-slate-600" : "bg-green-600"}`}>
                                <span>{todo.todo}</span>
                            </div>
                            <div className="flex flex-row my-2 py-2 mx-4 cursor-pointer">
                                {
                                    <div onClick={() => { updateTODODone(todo.todo_id, todo.is_done) }}>
                                        {
                                            todo.is_done ? <MdOutlineCheckBox size={25} /> : <MdCheckBoxOutlineBlank size={25} />
                                        }
                                    </div>
                                }
                                <div onClick={() => { deleteTODO(todo.todo_id) }}>
                                    <FaTrashCan size={25} className="ml-2" />
                                </div>
                            </div>
                        </div>
                    ))
                )
            }
        </div>
    )
}

export default TODOS;