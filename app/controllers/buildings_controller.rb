class BuildingsController < ApplicationController
  
  def index
    @items = Building.all.limit(50)
  end

end
