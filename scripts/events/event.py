#!/usr/bin/env python

import sys
import pycurl
import json
import signal
import time

stream_url = 'https://localhost:5665/v1/events?queue=test&types=CheckResult'
user = 'root'
password = 'icinga'

# {"check_result":{"active":true,"check_source":"mbmif.int.netways.de","command":null,"execution_end":1468583225.5626189709,"execution_start":1468583225.562431097,"exit_status":0.0,"output":"Hello from mbmif.int.netways.de","performance_data":[{"counter":false,"crit":null,"label":"time","max":null,"min":null,"type":"PerfdataValue","unit":"","value":1468583225.5625960827,"warn":null}],"schedule_end":1468583225.5626189709,"schedule_start":1468583225.5573730469,"state":1.0,"type":"CheckResult","vars_after":{"attempt":1.0,"reachable":false,"state":1.0,"state_type":1.0},"vars_before":{"attempt":1.0,"reachable":false,"state":3.0,"state_type":1.0}},"host":"icingacamp","service":"random-8","timestamp":1468583225.5629920959,"type":"CheckResult"}

class EventStream:
  debug = False

  def __init__(self):
    self.conn = None
    self.user = 'root'
    self.password = 'icinga'
    self.url = 'https://localhost:5665/v1/events'
    self.queue = 'es-test'
    #self.types = [ 'CheckResult', 'StateChange' ]
    self.types = [ 'CheckResult' ]
    # code parts only handle this type for now
    self.handle_type = 'source_rate'
    self.buffer = ''
    self.setup_connection()
    self.cr_source_rate = {}
    self.start_time = time.time()
    signal.signal(signal.SIGINT, lambda x,y: sys.exit(0)) #handle crtl+c gracefully

  def setup_connection(self):
    if self.conn:
       self.conn.close()
       self.buffer = ''

    url = self.url + '?queue=%s&%s' % (self.queue, '&'.join(['types[]=%s' % x for x in self.types]))
    #print url

    self.conn = pycurl.Curl()
    self.conn.setopt(pycurl.USERPWD, "%s:%s" % (self.user, self.password))
    self.conn.setopt(pycurl.URL, url)
    self.conn.setopt(pycurl.WRITEFUNCTION, self.handle_result)
    self.conn.setopt(pycurl.SSL_VERIFYPEER, 0)
    self.conn.setopt(pycurl.SSL_VERIFYHOST, 0)
    self.conn.setopt(pycurl.POST, True)
    self.conn.setopt(pycurl.HTTPHEADER, [ "Accept: application/json" ])
    if self.debug:
      self.conn.setopt(pycurl.VERBOSE, True)

  def start(self):
    try:
      self.conn.perform()
    except pycurl.error, err:
      pass
      #print 'Network error: %s' % str(err)

    self.stop()

  def stop(self):
    self.conn.close()

  def handle_result(self, data):
    self.buffer += data
    try:
      result = json.loads(self.buffer)
      self.buffer = ''
    except ValueError:
      return

    # reserve multiple type handlers
    if "source_rate" in self.handle_type:
      cr = result["check_result"]

      if self.debug is True:
        print "Check Result: %s" % (json.dumps(cr))

      cs = cr["check_source"]
      if not cs in self.cr_source_rate.keys():
        self.cr_source_rate[cs] = 1
      else:
        self.cr_source_rate[cs] += 1

      time_elapsed = time.time() - self.start_time

      for source, count in self.cr_source_rate.iteritems():
         print "Check source average rate for '%s': %.2f/s" % (source, float(count/time_elapsed))


if __name__ == "__main__":
  es = EventStream()
  #es.debug = True
  es.debug = False
  es.setup_connection()
  try:
    es.start()
  except Exception:
    pass

  sys.exit(0)
