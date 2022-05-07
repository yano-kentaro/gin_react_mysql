import React from 'react'
import FullCalendar, { EventApi,
    DateSelectArg,EventClickArg,
    EventContentArg, formatDate, EventDropArg
  } from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin, { EventResizeDoneArg } from '@fullcalendar/interaction'
import { INITIAL_EVENTS, createEventId } from './event-utils'
import axios from 'axios'

interface DemoAppState {
  weekendsVisible: boolean
  currentEvents: EventApi[]
}

export default class DemoApp extends React.Component<{}, DemoAppState> {

  state: DemoAppState = {
    weekendsVisible: true,
    currentEvents: []
  }

  render() {
    return (
      <div className='demo-app'>
        {/* {this.renderSidebar()} */}
        <div className='demo-app-main'>
          <FullCalendar
            plugins={[
              dayGridPlugin,
              timeGridPlugin,
              interactionPlugin
            ]}
            headerToolbar={{
              left: 'prev,next today',
              center: 'title',
              right: 'dayGridMonth,timeGridWeek,timeGridDay'
            }}
            events={{
              url: "/api/events"
            }}
            initialView='dayGridMonth'
            editable={true}
            eventResize={this.eventResize}
            eventDrop={this.eventDrop}
            selectable={true}
            selectMirror={true}
            dayMaxEvents={true}
            weekends={this.state.weekendsVisible}
            //initialEvents={INITIAL_EVENTS} // alternatively, use the `events` setting to fetch from a feed
            select={this.handleDateSelect}
            eventContent={renderEventContent} // custom render function
            eventClick={this.handleEventClick}
            //eventsSet={this.handleEvents} // called after events are initialized/added/changed/removed
            /* you can update a remote database when these fire:
            eventAdd={function(){}}
            eventChange={function(){}}
            eventRemove={function(){}}
            */
          />
        </div>
      </div>
    )
  }

  // renderSidebar() {
  //   return (
  //     <div className='demo-app-sidebar'>
  //       <div className='demo-app-sidebar-section'>
  //         <h2>Instructions</h2>
  //         <ul>
  //           <li>Select dates and you will be prompted to create a new event</li>
  //           <li>Drag, drop, and resize events</li>
  //           <li>Click an event to delete it</li>
  //         </ul>
  //       </div>
  //       <div className='demo-app-sidebar-section'>
  //         <label>
  //           <input
  //             type='checkbox'
  //             checked={this.state.weekendsVisible}
  //             onChange={this.handleWeekendsToggle}
  //           ></input>
  //           toggle weekends
  //         </label>
  //       </div>
  //       <div className='demo-app-sidebar-section'>
  //         <h2>All Events ({this.state.currentEvents.length})</h2>
  //         <ul>
  //           {this.state.currentEvents.map(renderSidebarEvent)}
  //         </ul>
  //       </div>
  //     </div>
  //   )
  // }

  handleWeekendsToggle = () => {
    this.setState({
      weekendsVisible: !this.state.weekendsVisible
    })
  }

  // 新規登録
  handleDateSelect = async (selectInfo: DateSelectArg) => {
    let title = prompt('Please enter a new title for your event')
    let calendarApi = selectInfo.view.calendar

    calendarApi.unselect() // clear date selection

    if(title) {
      const res = await axios.post('/api/event/create', {
        title: title,
        start: selectInfo.startStr,
        end: selectInfo.endStr,
      })
      calendarApi.refetchEvents()
    } else {
      alert("title is required.")
    }
  }

  // 予定期間変更
  eventResize = async (resizeInfo: EventResizeDoneArg) => {
    let calendarApi = resizeInfo.view.calendar
    let newEvent = resizeInfo.event
    const res = await axios.put('/api/event/update/' + newEvent.id, {
      title: newEvent.title,
      start: newEvent.startStr,
      end: newEvent.endStr
    })
    calendarApi.refetchEvents()
  }

  // 予定開始時刻変更
  eventDrop = async (dropInfo: EventDropArg) => {
    let calendarApi = dropInfo.view.calendar
    let newEvent = dropInfo.event
    const res = await axios.put('/api/event/update/' + newEvent.id, {
      title: newEvent.title,
      start: newEvent.startStr,
      end: newEvent.endStr
    })
    calendarApi.refetchEvents()
  }

  // 削除
  handleEventClick = async (clickInfo: EventClickArg) => {
    if (confirm(`Are you sure you want to delete the event '${clickInfo.event.title}'`)) {
      let calendarApi = clickInfo.view.calendar
      const res = await axios.delete('/api/event/delete/' + clickInfo.event.id)
      calendarApi.refetchEvents()
    }
  }

  handleEvents = (events: EventApi[]) => {
    this.setState({
      currentEvents: events
    })
  }

}

function renderEventContent(eventContent: EventContentArg) {
  return (
    <>
      <b>{eventContent.timeText}</b>
      <i>{eventContent.event.title}</i>
    </>
  )
}

// function renderSidebarEvent(event: EventApi) {
//   return (
//     <li key={event.id}>
//       <b>{formatDate(event.start!, {year: 'numeric', month: 'short', day: 'numeric'})}</b>
//       <i>{event.title}</i>
//     </li>
//   )
// }
